// Copyright 2021 Dolthub, Inc.
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

package index

import (
	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/dolt/go/libraries/doltcore/schema"
	"github.com/dolthub/dolt/go/store/prolly"
	"github.com/dolthub/dolt/go/store/prolly/tree"
	"github.com/dolthub/dolt/go/store/val"
)

type prollyRowIter struct {
	iter prolly.MapIter
	ns   tree.NodeStore

	keyDesc val.TupleDesc
	valDesc val.TupleDesc

	keyProj []int
	valProj []int
	// ordProj is a concatenated list of output ordinals for |keyProj| and |valProj|
	ordProj []int
	rowLen  int
}

var _ sql.RowIter = prollyRowIter{}

func NewProllyRowIter(sch schema.Schema, rows prolly.Map, iter prolly.MapIter, projections []uint64) (sql.RowIter, error) {
	if projections == nil {
		projections = sch.GetAllCols().Tags
	}

	keyProj, valProj, ordProj := projectionMappings(sch, projections)
	kd, vd := rows.Descriptors()

	if schema.IsKeyless(sch) {
		return &prollyKeylessIter{
			iter:    iter,
			valDesc: vd,
			valProj: valProj,
			ordProj: ordProj,
			rowLen:  len(projections),
			ns:      rows.NodeStore(),
		}, nil
	}

	return prollyRowIter{
		iter:    iter,
		keyDesc: kd,
		valDesc: vd,
		keyProj: keyProj,
		valProj: valProj,
		ordProj: ordProj,
		rowLen:  len(projections),
		ns:      rows.NodeStore(),
	}, nil
}

// projectionMappings returns data structures that specify 1) which fields we read
// from key and value tuples, and 2) the position of those fields in the output row.
func projectionMappings(sch schema.Schema, projections []uint64) (keyMap, valMap, ordMap val.OrdinalMapping) {
	pks := sch.GetPKCols()
	nonPks := sch.GetNonPKCols()

	// Our mappings should only contain the physical columns of the schema
	numPhysicalColumns := len(projections)
	if schema.IsVirtual(sch) {
		numPhysicalColumns = 0
		for _, t := range projections {
			if idx, ok := pks.TagToIdx[t]; ok && !sch.GetAllCols().GetByIndex(idx).Virtual {
				numPhysicalColumns++
			} else if idx, ok := nonPks.TagToIdx[t]; ok && !sch.GetAllCols().GetByIndex(idx).Virtual {
				numPhysicalColumns++
			}
		}
	}
	
	allMap := make([]int, 2*numPhysicalColumns)
	i := 0
	j := numPhysicalColumns - 1
	for k, t := range projections {
		if idx, ok := pks.TagToIdx[t]; ok && !pks.GetByIndex(idx).Virtual {
			allMap[numPhysicalColumns+i] = k
			allMap[i] = idx
			i++
		} else if idx, ok := nonPks.TagToIdx[t]; ok && !nonPks.GetByIndex(idx).Virtual {
			allMap[j] = idx
			allMap[numPhysicalColumns+j] = k
			j--
		}
	}
	
	keyMap = allMap[:i]
	valMap = allMap[i:numPhysicalColumns]
	ordMap = allMap[numPhysicalColumns:]
	if schema.IsKeyless(sch) {
		// skip the cardinality value, increment every index
		for i := range keyMap {
			keyMap[i]++
		}
		for i := range valMap {
			valMap[i]++
		}
	}
	
	return
}

func (it prollyRowIter) Next(ctx *sql.Context) (sql.Row, error) {
	key, value, err := it.iter.Next(ctx)
	if err != nil {
		return nil, err
	}

	row := make(sql.Row, it.rowLen)
	for i, idx := range it.keyProj {
		outputIdx := it.ordProj[i]
		row[outputIdx], err = GetField(ctx, it.keyDesc, idx, key, it.ns)
		if err != nil {
			return nil, err
		}
	}
	for i, idx := range it.valProj {
		outputIdx := it.ordProj[len(it.keyProj)+i]
		row[outputIdx], err = GetField(ctx, it.valDesc, idx, value, it.ns)
		if err != nil {
			return nil, err
		}
	}
	return row, nil
}

func (it prollyRowIter) Close(ctx *sql.Context) error {
	return nil
}

type prollyKeylessIter struct {
	iter prolly.MapIter
	ns   tree.NodeStore

	valDesc val.TupleDesc
	valProj []int
	ordProj []int
	rowLen  int

	curr sql.Row
	card uint64
}

var _ sql.RowIter = &prollyKeylessIter{}

//var _ sql.RowIter2 = prollyKeylessIter{}

func (it *prollyKeylessIter) Next(ctx *sql.Context) (sql.Row, error) {
	if it.card == 0 {
		if err := it.nextTuple(ctx); err != nil {
			return nil, err
		}
	}

	it.card--

	return it.curr, nil
}

func (it *prollyKeylessIter) nextTuple(ctx *sql.Context) error {
	_, value, err := it.iter.Next(ctx)
	if err != nil {
		return err
	}

	it.card = val.ReadKeylessCardinality(value)
	it.curr = make(sql.Row, it.rowLen)

	for i, idx := range it.valProj {
		outputIdx := it.ordProj[i]
		it.curr[outputIdx], err = GetField(ctx, it.valDesc, idx, value, it.ns)
		if err != nil {
			return err
		}
	}
	return nil
}

func (it *prollyKeylessIter) Close(ctx *sql.Context) error {
	return nil
}
