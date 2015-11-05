// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_with_list_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("StructWithList",
			[]types.Field{
				types.Field{"l", types.MakeCompoundTypeRef(types.ListKind, types.MakePrimitiveTypeRef(types.UInt8Kind)), false},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), false},
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"i", types.MakePrimitiveTypeRef(types.Int64Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_with_list_CachedRef = types.RegisterPackage(&p)
}

// StructWithList

type StructWithList struct {
	_l ListOfUInt8
	_b bool
	_s string
	_i int64

	ref *ref.Ref
}

func NewStructWithList() StructWithList {
	return StructWithList{
		_l: NewListOfUInt8(),
		_b: false,
		_s: "",
		_i: int64(0),

		ref: &ref.Ref{},
	}
}

type StructWithListDef struct {
	L ListOfUInt8Def
	B bool
	S string
	I int64
}

func (def StructWithListDef) New() StructWithList {
	return StructWithList{
		_l:  def.L.New(),
		_b:  def.B,
		_s:  def.S,
		_i:  def.I,
		ref: &ref.Ref{},
	}
}

func (s StructWithList) Def() (d StructWithListDef) {
	d.L = s._l.Def()
	d.B = s._b
	d.S = s._s
	d.I = s._i
	return
}

var __typeRefForStructWithList types.TypeRef

func (m StructWithList) TypeRef() types.TypeRef {
	return __typeRefForStructWithList
}

func init() {
	__typeRefForStructWithList = types.MakeTypeRef(__genPackageInFile_struct_with_list_CachedRef, 0)
	types.RegisterStruct(__typeRefForStructWithList, builderForStructWithList, readerForStructWithList)
}

func builderForStructWithList(values []types.Value) types.Value {
	i := 0
	s := StructWithList{ref: &ref.Ref{}}
	s._l = values[i].(ListOfUInt8)
	i++
	s._b = bool(values[i].(types.Bool))
	i++
	s._s = values[i].(types.String).String()
	i++
	s._i = int64(values[i].(types.Int64))
	i++
	return s
}

func readerForStructWithList(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(StructWithList)
	values = append(values, s._l)
	values = append(values, types.Bool(s._b))
	values = append(values, types.NewString(s._s))
	values = append(values, types.Int64(s._i))
	return values
}

func (s StructWithList) Equals(other types.Value) bool {
	return other != nil && __typeRefForStructWithList.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s StructWithList) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructWithList) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForStructWithList.Chunks()...)
	chunks = append(chunks, s._l.Chunks()...)
	return
}

func (s StructWithList) ChildValues() (ret []types.Value) {
	ret = append(ret, s._l)
	ret = append(ret, types.Bool(s._b))
	ret = append(ret, types.NewString(s._s))
	ret = append(ret, types.Int64(s._i))
	return
}

func (s StructWithList) L() ListOfUInt8 {
	return s._l
}

func (s StructWithList) SetL(val ListOfUInt8) StructWithList {
	s._l = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithList) B() bool {
	return s._b
}

func (s StructWithList) SetB(val bool) StructWithList {
	s._b = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithList) S() string {
	return s._s
}

func (s StructWithList) SetS(val string) StructWithList {
	s._s = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithList) I() int64 {
	return s._i
}

func (s StructWithList) SetI(val int64) StructWithList {
	s._i = val
	s.ref = &ref.Ref{}
	return s
}

// ListOfUInt8

type ListOfUInt8 struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfUInt8() ListOfUInt8 {
	return ListOfUInt8{types.NewList(), &ref.Ref{}}
}

type ListOfUInt8Def []uint8

func (def ListOfUInt8Def) New() ListOfUInt8 {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = types.UInt8(d)
	}
	return ListOfUInt8{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfUInt8) Def() ListOfUInt8Def {
	d := make([]uint8, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = uint8(l.l.Get(i).(types.UInt8))
	}
	return d
}

func (l ListOfUInt8) Equals(other types.Value) bool {
	return other != nil && __typeRefForListOfUInt8.Equals(other.TypeRef()) && l.Ref() == other.Ref()
}

func (l ListOfUInt8) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfUInt8) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.TypeRef().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfUInt8) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfUInt8.
var __typeRefForListOfUInt8 types.TypeRef

func (m ListOfUInt8) TypeRef() types.TypeRef {
	return __typeRefForListOfUInt8
}

func init() {
	__typeRefForListOfUInt8 = types.MakeCompoundTypeRef(types.ListKind, types.MakePrimitiveTypeRef(types.UInt8Kind))
	types.RegisterValue(__typeRefForListOfUInt8, builderForListOfUInt8, readerForListOfUInt8)
}

func builderForListOfUInt8(v types.Value) types.Value {
	return ListOfUInt8{v.(types.List), &ref.Ref{}}
}

func readerForListOfUInt8(v types.Value) types.Value {
	return v.(ListOfUInt8).l
}

func (l ListOfUInt8) Len() uint64 {
	return l.l.Len()
}

func (l ListOfUInt8) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfUInt8) Get(i uint64) uint8 {
	return uint8(l.l.Get(i).(types.UInt8))
}

func (l ListOfUInt8) Slice(idx uint64, end uint64) ListOfUInt8 {
	return ListOfUInt8{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfUInt8) Set(i uint64, val uint8) ListOfUInt8 {
	return ListOfUInt8{l.l.Set(i, types.UInt8(val)), &ref.Ref{}}
}

func (l ListOfUInt8) Append(v ...uint8) ListOfUInt8 {
	return ListOfUInt8{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfUInt8) Insert(idx uint64, v ...uint8) ListOfUInt8 {
	return ListOfUInt8{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfUInt8) Remove(idx uint64, end uint64) ListOfUInt8 {
	return ListOfUInt8{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfUInt8) RemoveAt(idx uint64) ListOfUInt8 {
	return ListOfUInt8{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfUInt8) fromElemSlice(p []uint8) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.UInt8(v)
	}
	return r
}

type ListOfUInt8IterCallback func(v uint8, i uint64) (stop bool)

func (l ListOfUInt8) Iter(cb ListOfUInt8IterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(uint8(v.(types.UInt8)), i)
	})
}

type ListOfUInt8IterAllCallback func(v uint8, i uint64)

func (l ListOfUInt8) IterAll(cb ListOfUInt8IterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(uint8(v.(types.UInt8)), i)
	})
}

func (l ListOfUInt8) IterAllP(concurrency int, cb ListOfUInt8IterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(uint8(v.(types.UInt8)), i)
	})
}

type ListOfUInt8FilterCallback func(v uint8, i uint64) (keep bool)

func (l ListOfUInt8) Filter(cb ListOfUInt8FilterCallback) ListOfUInt8 {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(uint8(v.(types.UInt8)), i)
	})
	return ListOfUInt8{out, &ref.Ref{}}
}
