// This file was generated by nomdl/codegen.

package types

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
)

var __typesPackageInFile_compound_blob_struct_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := NewPackage([]TypeRef{
		MakeStructTypeRef("compoundBlobStruct",
			[]Field{
				Field{"Offsets", MakeCompoundTypeRef(ListKind, MakePrimitiveTypeRef(UInt64Kind)), false},
				Field{"Blobs", MakeCompoundTypeRef(ListKind, MakeCompoundTypeRef(RefKind, MakePrimitiveTypeRef(BlobKind))), false},
			},
			Choices{},
		),
	}, []ref.Ref{})
	__typesPackageInFile_compound_blob_struct_CachedRef = RegisterPackage(&p)
}

// compoundBlobStruct

type compoundBlobStruct struct {
	_Offsets ListOfUInt64
	_Blobs   ListOfRefOfBlob

	ref *ref.Ref
}

func NewcompoundBlobStruct() compoundBlobStruct {
	return compoundBlobStruct{
		_Offsets: NewListOfUInt64(),
		_Blobs:   NewListOfRefOfBlob(),

		ref: &ref.Ref{},
	}
}

type compoundBlobStructDef struct {
	Offsets ListOfUInt64Def
	Blobs   ListOfRefOfBlobDef
}

func (def compoundBlobStructDef) New() compoundBlobStruct {
	return compoundBlobStruct{
		_Offsets: def.Offsets.New(),
		_Blobs:   def.Blobs.New(),
		ref:      &ref.Ref{},
	}
}

func (s compoundBlobStruct) Def() (d compoundBlobStructDef) {
	d.Offsets = s._Offsets.Def()
	d.Blobs = s._Blobs.Def()
	return
}

var __typeRefForcompoundBlobStruct TypeRef

func (m compoundBlobStruct) TypeRef() TypeRef {
	return __typeRefForcompoundBlobStruct
}

func init() {
	__typeRefForcompoundBlobStruct = MakeTypeRef(__typesPackageInFile_compound_blob_struct_CachedRef, 0)
	RegisterStruct(__typeRefForcompoundBlobStruct, builderForcompoundBlobStruct, readerForcompoundBlobStruct)
}

func builderForcompoundBlobStruct(values []Value) Value {
	i := 0
	s := compoundBlobStruct{ref: &ref.Ref{}}
	s._Offsets = values[i].(ListOfUInt64)
	i++
	s._Blobs = values[i].(ListOfRefOfBlob)
	i++
	return s
}

func readerForcompoundBlobStruct(v Value) []Value {
	values := []Value{}
	s := v.(compoundBlobStruct)
	values = append(values, s._Offsets)
	values = append(values, s._Blobs)
	return values
}

func (s compoundBlobStruct) Equals(other Value) bool {
	return other != nil && __typeRefForcompoundBlobStruct.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s compoundBlobStruct) Ref() ref.Ref {
	return EnsureRef(s.ref, s)
}

func (s compoundBlobStruct) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForcompoundBlobStruct.Chunks()...)
	chunks = append(chunks, s._Offsets.Chunks()...)
	chunks = append(chunks, s._Blobs.Chunks()...)
	return
}

func (s compoundBlobStruct) ChildValues() (ret []Value) {
	ret = append(ret, s._Offsets)
	ret = append(ret, s._Blobs)
	return
}

func (s compoundBlobStruct) Offsets() ListOfUInt64 {
	return s._Offsets
}

func (s compoundBlobStruct) SetOffsets(val ListOfUInt64) compoundBlobStruct {
	s._Offsets = val
	s.ref = &ref.Ref{}
	return s
}

func (s compoundBlobStruct) Blobs() ListOfRefOfBlob {
	return s._Blobs
}

func (s compoundBlobStruct) SetBlobs(val ListOfRefOfBlob) compoundBlobStruct {
	s._Blobs = val
	s.ref = &ref.Ref{}
	return s
}

// ListOfUInt64

type ListOfUInt64 struct {
	l   List
	ref *ref.Ref
}

func NewListOfUInt64() ListOfUInt64 {
	return ListOfUInt64{NewList(), &ref.Ref{}}
}

type ListOfUInt64Def []uint64

func (def ListOfUInt64Def) New() ListOfUInt64 {
	l := make([]Value, len(def))
	for i, d := range def {
		l[i] = UInt64(d)
	}
	return ListOfUInt64{NewList(l...), &ref.Ref{}}
}

func (l ListOfUInt64) Def() ListOfUInt64Def {
	d := make([]uint64, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = uint64(l.l.Get(i).(UInt64))
	}
	return d
}

func (l ListOfUInt64) Equals(other Value) bool {
	return other != nil && __typeRefForListOfUInt64.Equals(other.TypeRef()) && l.Ref() == other.Ref()
}

func (l ListOfUInt64) Ref() ref.Ref {
	return EnsureRef(l.ref, l)
}

func (l ListOfUInt64) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.TypeRef().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfUInt64) ChildValues() []Value {
	return append([]Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfUInt64.
var __typeRefForListOfUInt64 TypeRef

func (m ListOfUInt64) TypeRef() TypeRef {
	return __typeRefForListOfUInt64
}

func init() {
	__typeRefForListOfUInt64 = MakeCompoundTypeRef(ListKind, MakePrimitiveTypeRef(UInt64Kind))
	RegisterValue(__typeRefForListOfUInt64, builderForListOfUInt64, readerForListOfUInt64)
}

func builderForListOfUInt64(v Value) Value {
	return ListOfUInt64{v.(List), &ref.Ref{}}
}

func readerForListOfUInt64(v Value) Value {
	return v.(ListOfUInt64).l
}

func (l ListOfUInt64) Len() uint64 {
	return l.l.Len()
}

func (l ListOfUInt64) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfUInt64) Get(i uint64) uint64 {
	return uint64(l.l.Get(i).(UInt64))
}

func (l ListOfUInt64) Slice(idx uint64, end uint64) ListOfUInt64 {
	return ListOfUInt64{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfUInt64) Set(i uint64, val uint64) ListOfUInt64 {
	return ListOfUInt64{l.l.Set(i, UInt64(val)), &ref.Ref{}}
}

func (l ListOfUInt64) Append(v ...uint64) ListOfUInt64 {
	return ListOfUInt64{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfUInt64) Insert(idx uint64, v ...uint64) ListOfUInt64 {
	return ListOfUInt64{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfUInt64) Remove(idx uint64, end uint64) ListOfUInt64 {
	return ListOfUInt64{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfUInt64) RemoveAt(idx uint64) ListOfUInt64 {
	return ListOfUInt64{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfUInt64) fromElemSlice(p []uint64) []Value {
	r := make([]Value, len(p))
	for i, v := range p {
		r[i] = UInt64(v)
	}
	return r
}

type ListOfUInt64IterCallback func(v uint64, i uint64) (stop bool)

func (l ListOfUInt64) Iter(cb ListOfUInt64IterCallback) {
	l.l.Iter(func(v Value, i uint64) bool {
		return cb(uint64(v.(UInt64)), i)
	})
}

type ListOfUInt64IterAllCallback func(v uint64, i uint64)

func (l ListOfUInt64) IterAll(cb ListOfUInt64IterAllCallback) {
	l.l.IterAll(func(v Value, i uint64) {
		cb(uint64(v.(UInt64)), i)
	})
}

func (l ListOfUInt64) IterAllP(concurrency int, cb ListOfUInt64IterAllCallback) {
	l.l.IterAllP(concurrency, func(v Value, i uint64) {
		cb(uint64(v.(UInt64)), i)
	})
}

type ListOfUInt64FilterCallback func(v uint64, i uint64) (keep bool)

func (l ListOfUInt64) Filter(cb ListOfUInt64FilterCallback) ListOfUInt64 {
	out := l.l.Filter(func(v Value, i uint64) bool {
		return cb(uint64(v.(UInt64)), i)
	})
	return ListOfUInt64{out, &ref.Ref{}}
}

// ListOfRefOfBlob

type ListOfRefOfBlob struct {
	l   List
	ref *ref.Ref
}

func NewListOfRefOfBlob() ListOfRefOfBlob {
	return ListOfRefOfBlob{NewList(), &ref.Ref{}}
}

type ListOfRefOfBlobDef []ref.Ref

func (def ListOfRefOfBlobDef) New() ListOfRefOfBlob {
	l := make([]Value, len(def))
	for i, d := range def {
		l[i] = NewRefOfBlob(d)
	}
	return ListOfRefOfBlob{NewList(l...), &ref.Ref{}}
}

func (l ListOfRefOfBlob) Def() ListOfRefOfBlobDef {
	d := make([]ref.Ref, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(RefOfBlob).TargetRef()
	}
	return d
}

func (l ListOfRefOfBlob) Equals(other Value) bool {
	return other != nil && __typeRefForListOfRefOfBlob.Equals(other.TypeRef()) && l.Ref() == other.Ref()
}

func (l ListOfRefOfBlob) Ref() ref.Ref {
	return EnsureRef(l.ref, l)
}

func (l ListOfRefOfBlob) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.TypeRef().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfRefOfBlob) ChildValues() []Value {
	return append([]Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfRefOfBlob.
var __typeRefForListOfRefOfBlob TypeRef

func (m ListOfRefOfBlob) TypeRef() TypeRef {
	return __typeRefForListOfRefOfBlob
}

func init() {
	__typeRefForListOfRefOfBlob = MakeCompoundTypeRef(ListKind, MakeCompoundTypeRef(RefKind, MakePrimitiveTypeRef(BlobKind)))
	RegisterValue(__typeRefForListOfRefOfBlob, builderForListOfRefOfBlob, readerForListOfRefOfBlob)
}

func builderForListOfRefOfBlob(v Value) Value {
	return ListOfRefOfBlob{v.(List), &ref.Ref{}}
}

func readerForListOfRefOfBlob(v Value) Value {
	return v.(ListOfRefOfBlob).l
}

func (l ListOfRefOfBlob) Len() uint64 {
	return l.l.Len()
}

func (l ListOfRefOfBlob) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfRefOfBlob) Get(i uint64) RefOfBlob {
	return l.l.Get(i).(RefOfBlob)
}

func (l ListOfRefOfBlob) Slice(idx uint64, end uint64) ListOfRefOfBlob {
	return ListOfRefOfBlob{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfBlob) Set(i uint64, val RefOfBlob) ListOfRefOfBlob {
	return ListOfRefOfBlob{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfRefOfBlob) Append(v ...RefOfBlob) ListOfRefOfBlob {
	return ListOfRefOfBlob{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfBlob) Insert(idx uint64, v ...RefOfBlob) ListOfRefOfBlob {
	return ListOfRefOfBlob{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfBlob) Remove(idx uint64, end uint64) ListOfRefOfBlob {
	return ListOfRefOfBlob{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfBlob) RemoveAt(idx uint64) ListOfRefOfBlob {
	return ListOfRefOfBlob{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfRefOfBlob) fromElemSlice(p []RefOfBlob) []Value {
	r := make([]Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfRefOfBlobIterCallback func(v RefOfBlob, i uint64) (stop bool)

func (l ListOfRefOfBlob) Iter(cb ListOfRefOfBlobIterCallback) {
	l.l.Iter(func(v Value, i uint64) bool {
		return cb(v.(RefOfBlob), i)
	})
}

type ListOfRefOfBlobIterAllCallback func(v RefOfBlob, i uint64)

func (l ListOfRefOfBlob) IterAll(cb ListOfRefOfBlobIterAllCallback) {
	l.l.IterAll(func(v Value, i uint64) {
		cb(v.(RefOfBlob), i)
	})
}

func (l ListOfRefOfBlob) IterAllP(concurrency int, cb ListOfRefOfBlobIterAllCallback) {
	l.l.IterAllP(concurrency, func(v Value, i uint64) {
		cb(v.(RefOfBlob), i)
	})
}

type ListOfRefOfBlobFilterCallback func(v RefOfBlob, i uint64) (keep bool)

func (l ListOfRefOfBlob) Filter(cb ListOfRefOfBlobFilterCallback) ListOfRefOfBlob {
	out := l.l.Filter(func(v Value, i uint64) bool {
		return cb(v.(RefOfBlob), i)
	})
	return ListOfRefOfBlob{out, &ref.Ref{}}
}

// RefOfBlob

type RefOfBlob struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfBlob(target ref.Ref) RefOfBlob {
	return RefOfBlob{target, &ref.Ref{}}
}

func (r RefOfBlob) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfBlob) Ref() ref.Ref {
	return EnsureRef(r.ref, r)
}

func (r RefOfBlob) Equals(other Value) bool {
	return other != nil && __typeRefForRefOfBlob.Equals(other.TypeRef()) && r.Ref() == other.Ref()
}

func (r RefOfBlob) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.TypeRef().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfBlob) ChildValues() []Value {
	return nil
}

// A Noms Value that describes RefOfBlob.
var __typeRefForRefOfBlob TypeRef

func (m RefOfBlob) TypeRef() TypeRef {
	return __typeRefForRefOfBlob
}

func init() {
	__typeRefForRefOfBlob = MakeCompoundTypeRef(RefKind, MakePrimitiveTypeRef(BlobKind))
	RegisterRef(__typeRefForRefOfBlob, builderForRefOfBlob)
}

func builderForRefOfBlob(r ref.Ref) Value {
	return NewRefOfBlob(r)
}

func (r RefOfBlob) TargetValue(cs chunks.ChunkSource) Blob {
	return ReadValue(r.target, cs).(Blob)
}

func (r RefOfBlob) SetTargetValue(val Blob, cs chunks.ChunkSink) RefOfBlob {
	return NewRefOfBlob(WriteValue(val, cs))
}
