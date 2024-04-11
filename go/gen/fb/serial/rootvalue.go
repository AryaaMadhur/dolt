// Copyright 2022-2023 Dolthub, Inc.
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

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serial

import (
	flatbuffers "github.com/dolthub/flatbuffers/v23/go"
)

type RootValue struct {
	_tab flatbuffers.Table
}

func InitRootValueRoot(o *RootValue, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	return o.Init(buf, n+offset)
}

func TryGetRootAsRootValue(buf []byte, offset flatbuffers.UOffsetT) (*RootValue, error) {
	x := &RootValue{}
	return x, InitRootValueRoot(x, buf, offset)
}

func TryGetSizePrefixedRootAsRootValue(buf []byte, offset flatbuffers.UOffsetT) (*RootValue, error) {
	x := &RootValue{}
	return x, InitRootValueRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func (rcv *RootValue) Init(buf []byte, i flatbuffers.UOffsetT) error {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
	if RootValueNumFields < rcv.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func (rcv *RootValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RootValue) FeatureVersion() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RootValue) MutateFeatureVersion(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *RootValue) Tables(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *RootValue) TablesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RootValue) TablesBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RootValue) MutateTables(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *RootValue) ForeignKeyAddr(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *RootValue) ForeignKeyAddrLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RootValue) ForeignKeyAddrBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RootValue) MutateForeignKeyAddr(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *RootValue) Collation() Collation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return Collation(rcv._tab.GetUint16(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *RootValue) MutateCollation(n Collation) bool {
	return rcv._tab.MutateUint16Slot(10, uint16(n))
}

func (rcv *RootValue) TrySchemas(obj *DatabaseSchema, j int) (bool, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		if DatabaseSchemaNumFields < obj.Table().NumFields() {
			return false, flatbuffers.ErrTableHasUnknownFields
		}
		return true, nil
	}
	return false, nil
}

func (rcv *RootValue) SchemasLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

const RootValueNumFields = 5

func RootValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(RootValueNumFields)
}
func RootValueAddFeatureVersion(builder *flatbuffers.Builder, featureVersion int64) {
	builder.PrependInt64Slot(0, featureVersion, 0)
}
func RootValueAddTables(builder *flatbuffers.Builder, tables flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(tables), 0)
}
func RootValueStartTablesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func RootValueAddForeignKeyAddr(builder *flatbuffers.Builder, foreignKeyAddr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(foreignKeyAddr), 0)
}
func RootValueStartForeignKeyAddrVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func RootValueAddCollation(builder *flatbuffers.Builder, collation Collation) {
	builder.PrependUint16Slot(3, uint16(collation), 0)
}
func RootValueAddSchemas(builder *flatbuffers.Builder, schemas flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(schemas), 0)
}
func RootValueStartSchemasVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RootValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
