// Copyright 2019 Dolthub, Inc.
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

package actions

import (
	"bytes"
	"strings"

	"github.com/dolthub/dolt/go/libraries/doltcore/diff"
)

type tblErrorType string

const (
	tblErrInvalid         tblErrorType = "invalid"
	tblErrTypeNotExist    tblErrorType = "do not exist"
	tblErrTypeInConflict  tblErrorType = "are in conflict"
	tblErrTypeSchConflict tblErrorType = "have schema conflicts"
	tblErrTypeConstViols  tblErrorType = "have constraint violations"
)

type TblError struct {
	tables     []string
	tblErrType tblErrorType
}

func NewTblNotExistError(tbls []string) TblError {
	return TblError{tbls, tblErrTypeNotExist}
}

func NewTblInConflictError(tbls []string) TblError {
	return TblError{tbls, tblErrTypeInConflict}
}

func NewTblSchemaConflictError(tbls []string) TblError {
	return TblError{tbls, tblErrTypeSchConflict}
}

func NewTblHasConstraintViolations(tbls []string) TblError {
	return TblError{tbls, tblErrTypeConstViols}
}

func (te TblError) Error() string {
	return "error: the table(s) " + strings.Join(te.tables, ", ") + " " + string(te.tblErrType)
}

func getTblErrType(err error) tblErrorType {
	te, ok := err.(TblError)

	if ok {
		return te.tblErrType
	}

	return tblErrInvalid
}

func IsTblError(err error) bool {
	return getTblErrType(err) != tblErrInvalid
}

func IsTblNotExist(err error) bool {
	return getTblErrType(err) == tblErrTypeNotExist
}

func IsTblInConflict(err error) bool {
	return getTblErrType(err) == tblErrTypeInConflict
}

func IsTblViolatesConstraints(err error) bool {
	return getTblErrType(err) == tblErrTypeConstViols
}

func GetTablesForError(err error) []string {
	te, ok := err.(TblError)

	if !ok {
		panic("Must validate with IsTblError or more specific methods before calling GetTablesForError")
	}

	return te.tables
}

type ErrCheckoutWouldOverwrite struct {
	tables []string
}

func (cwo ErrCheckoutWouldOverwrite) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString("Your local changes to the following tables would be overwritten by checkout:\n")
	for _, tbl := range cwo.tables {
		buffer.WriteString("\t" + tbl + "\n")
	}

	buffer.WriteString("Please commit your changes or stash them before you switch branches.\n")
	buffer.WriteString("Aborting")
	return buffer.String()
}

func IsCheckoutWouldOverwrite(err error) bool {
	_, ok := err.(ErrCheckoutWouldOverwrite)
	return ok
}

func CheckoutWouldOverwriteTables(err error) []string {
	cwo, ok := err.(ErrCheckoutWouldOverwrite)

	if !ok {
		panic("Must validate with IsCheckoutWouldOverwrite before calling CheckoutWouldOverwriteTables")
	}

	return cwo.tables
}

type NothingStaged struct {
	NotStagedTbls []diff.TableDelta
}

func (ns NothingStaged) Error() string {
	return "no changes added to commit"
}

func IsNothingStaged(err error) bool {
	_, ok := err.(NothingStaged)
	return ok
}

func NothingStagedTblDiffs(err error) []diff.TableDelta {
	ns, ok := err.(NothingStaged)

	if !ok {
		panic("Must validate with IsCheckoutWouldOverwrite before calling CheckoutWouldOverwriteTables")
	}

	return ns.NotStagedTbls
}
