// Copyright 2019 Liquidata, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqle

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/src-d/go-mysql-server/sql"

	"github.com/liquidata-inc/dolt/go/cmd/dolt/errhand"
	"github.com/liquidata-inc/dolt/go/libraries/doltcore/doltdb"
	"github.com/liquidata-inc/dolt/go/libraries/doltcore/schema"
	"github.com/liquidata-inc/dolt/go/store/types"
)

// DoltTable implements the sql.Table interface and gives access to dolt table rows and schema.
type DoltTable struct {
	name  string
	table *doltdb.Table
	sch   schema.Schema
	db    *Database
}

// Implements sql.IndexableTable
func (t *DoltTable) WithIndexLookup(lookup sql.IndexLookup) sql.Table {
	dil, ok := lookup.(*doltIndexLookup)
	if !ok {
		panic(fmt.Sprintf("Unrecognized indexLookup %T", lookup))
	}

	return &IndexedDoltTable{
		table:       t,
		indexLookup: dil,
	}
}

// Implements sql.IndexableTable
func (t *DoltTable) IndexKeyValues(*sql.Context, []string) (sql.PartitionIndexKeyValueIter, error) {
	return nil, errors.New("creating new indexes not supported")
}

// Implements sql.IndexableTable
func (t *DoltTable) IndexLookup() sql.IndexLookup {
	panic("IndexLookup called on DoltTable, should be on IndexedDoltTable")
}

// Name returns the name of the table.
func (t *DoltTable) Name() string {
	return t.name
}

// Not sure what the purpose of this method is, so returning the name for now.
func (t *DoltTable) String() string {
	return t.name
}

// Schema returns the schema for this table.
func (t *DoltTable) Schema() sql.Schema {
	// TODO: fix panics
	sch, err := t.table.GetSchema(context.TODO())

	if err != nil {
		panic(err)
	}

	// TODO: fix panics
	sqlSch, err := doltSchemaToSqlSchema(t.name, sch)

	if err != nil {
		panic(err)
	}

	return sqlSch
}

// Returns the partitions for this table. We return a single partition, but could potentially get more performance by
// returning multiple.
func (t *DoltTable) Partitions(*sql.Context) (sql.PartitionIter, error) {
	return &doltTablePartitionIter{}, nil
}

// Returns the table rows for the partition given (all rows of the table).
func (t *DoltTable) PartitionRows(ctx *sql.Context, _ sql.Partition) (sql.RowIter, error) {
	return newRowIterator(t, ctx)
}

// Insert adds the given row to the table and updates the database root.
func (t *DoltTable) Insert(ctx *sql.Context, sqlRow sql.Row) error {
	dRow, err := SqlRowToDoltRow(t.table.Format(), sqlRow, t.sch)
	if err != nil {
		return err
	}

	key, err := dRow.NomsMapKey(t.sch).Value(ctx)
	if err != nil {
		return errhand.BuildDError("failed to get row key").AddCause(err).Build()
	}
	_, rowExists, err := t.table.GetRow(ctx, key.(types.Tuple), t.sch)
	if err != nil {
		return errhand.BuildDError("failed to read table").AddCause(err).Build()
	}
	if rowExists {
		return errors.New("duplicate primary key given")
	}

	return t.updateTable(ctx, dRow.NomsMapKey(t.sch), dRow.NomsMapValue(t.sch), nil)
}

// Delete removes the given row to the table and updates the database root.
func (t *DoltTable) Delete(ctx *sql.Context, sqlRow sql.Row) error {
	dRow, err := SqlRowToDoltRow(t.table.Format(), sqlRow, t.sch)
	if err != nil {
		return err
	}

	nomsMapKey := dRow.NomsMapKey(t.sch)
	nmkValue, err := nomsMapKey.Value(ctx)
	if err != nil {
		return errhand.BuildDError("failed to get row key").AddCause(err).Build()
	}
	_, rowExists, err := t.table.GetRow(ctx, nmkValue.(types.Tuple), t.sch)
	if err != nil {
		return errhand.BuildDError("failed to read table").AddCause(err).Build()
	}
	if !rowExists {
		return sql.ErrDeleteRowNotFound
	}

	return t.updateTable(ctx, nil, nil, nomsMapKey)
}

func (t *DoltTable) Update(ctx *sql.Context, oldRow sql.Row, newRow sql.Row) error {
	dOldRow, err := SqlRowToDoltRow(t.table.Format(), oldRow, t.sch)
	if err != nil {
		return err
	}
	dNewRow, err := SqlRowToDoltRow(t.table.Format(), newRow, t.sch)
	if err != nil {
		return err
	}

	// If the PK is changed then we have to delete the old row first
	// This is assuming that the new PK is not taken
	dOldKey := dOldRow.NomsMapKey(t.sch)
	dOldKeyVal, err := dOldKey.Value(ctx)
	if err != nil {
		return err
	}
	dNewKey := dNewRow.NomsMapKey(t.sch)
	dNewKeyVal, err := dNewKey.Value(ctx)
	if err != nil {
		return err
	}

	var keyToDelete types.LesserValuable
	if !dOldKeyVal.Equals(dNewKeyVal) {
		_, rowExists, err := t.table.GetRow(ctx, dNewKeyVal.(types.Tuple), t.sch)
		if err != nil {
			return errhand.BuildDError("failed to read table").AddCause(err).Build()
		}
		if rowExists {
			newRowAsStrings := make([]string, len(newRow))
			for i, val := range newRow {
				newRowAsStrings[i] = fmt.Sprintf("%v", val)
			}
			return fmt.Errorf("primary key collision: (%v)", strings.Join(newRowAsStrings, ", "))
		}
		keyToDelete = dOldKey
	}

	return t.updateTable(ctx, dNewRow.NomsMapKey(t.sch), dNewRow.NomsMapValue(t.sch), keyToDelete)
}

// doltTablePartitionIter, an object that knows how to return the single partition exactly once.
type doltTablePartitionIter struct {
	sql.PartitionIter
	i int
}

// Close is required by the sql.PartitionIter interface. Does nothing.
func (itr *doltTablePartitionIter) Close() error {
	return nil
}

// Next returns the next partition if there is one, or io.EOF if there isn't.
func (itr *doltTablePartitionIter) Next() (sql.Partition, error) {
	if itr.i > 0 {
		return nil, io.EOF
	}
	itr.i++

	return &doltTablePartition{}, nil
}

// A table partition, currently an unused layer of abstraction but required for the framework.
type doltTablePartition struct {
	sql.Partition
}

const partitionName = "single"

// Key returns the key for this partition, which must uniquely identity the partition. We have only a single partition
// per table, so we use a constant.
func (p doltTablePartition) Key() []byte {
	return []byte(partitionName)
}

// key & value are for the row to insert/update
// keyToDelete is to delete a row before insertion (such as updating a primary key)
func (t *DoltTable) updateTable(ctx *sql.Context, key types.LesserValuable, value types.Valuable, keyToDelete types.LesserValuable) error {
	typesMap, err := t.table.GetRowData(ctx)
	if err != nil {
		return errhand.BuildDError("failed to get row data.").AddCause(err).Build()
	}

	mapEditor := typesMap.Edit()
	if keyToDelete != nil {
		mapEditor = mapEditor.Remove(keyToDelete)
	}

	if key != nil && value != nil {
		mapEditor = mapEditor.Set(key, value)
	}

	updated, err := mapEditor.Map(ctx)
	if err != nil {
		return errhand.BuildDError("failed to modify table").AddCause(err).Build()
	}

	newTable, err := t.table.UpdateRows(ctx, updated)
	if err != nil {
		return errhand.BuildDError("failed to update rows").AddCause(err).Build()
	}

	newRoot, err := doltdb.PutTable(ctx, t.db.root, t.db.root.VRW(), t.name, newTable)
	if err != nil {
		return errhand.BuildDError("failed to write table back to database").AddCause(err).Build()
	}

	t.table = newTable
	t.db.root = newRoot
	return nil
}
