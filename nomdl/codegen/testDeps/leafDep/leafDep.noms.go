// This file was generated by nomdl/codegen.

package leafDep

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __leafDepPackageInFile_leafDep_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("S",
			[]types.Field{
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), false},
			},
			types.Choices{},
		),
		types.MakeEnumTypeRef("E", "e1", "e2", "e3"),
	}, []ref.Ref{})
	__leafDepPackageInFile_leafDep_CachedRef = types.RegisterPackage(&p)
}

// S

type S struct {
	_s string
	_b bool

	ref *ref.Ref
}

func NewS() S {
	return S{
		_s: "",
		_b: false,

		ref: &ref.Ref{},
	}
}

type SDef struct {
	S string
	B bool
}

func (def SDef) New() S {
	return S{
		_s:  def.S,
		_b:  def.B,
		ref: &ref.Ref{},
	}
}

func (s S) Def() (d SDef) {
	d.S = s._s
	d.B = s._b
	return
}

var __typeRefForS types.TypeRef

func (m S) TypeRef() types.TypeRef {
	return __typeRefForS
}

func init() {
	__typeRefForS = types.MakeTypeRef(__leafDepPackageInFile_leafDep_CachedRef, 0)
	types.RegisterStruct(__typeRefForS, builderForS, readerForS)
}

func builderForS(values []types.Value) types.Value {
	i := 0
	s := S{ref: &ref.Ref{}}
	s._s = values[i].(types.String).String()
	i++
	s._b = bool(values[i].(types.Bool))
	i++
	return s
}

func readerForS(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(S)
	values = append(values, types.NewString(s._s))
	values = append(values, types.Bool(s._b))
	return values
}

func (s S) Equals(other types.Value) bool {
	return other != nil && __typeRefForS.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s S) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s S) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForS.Chunks()...)
	return
}

func (s S) ChildValues() (ret []types.Value) {
	ret = append(ret, types.NewString(s._s))
	ret = append(ret, types.Bool(s._b))
	return
}

func (s S) S() string {
	return s._s
}

func (s S) SetS(val string) S {
	s._s = val
	s.ref = &ref.Ref{}
	return s
}

func (s S) B() bool {
	return s._b
}

func (s S) SetB(val bool) S {
	s._b = val
	s.ref = &ref.Ref{}
	return s
}

// E

type E uint32

const (
	E1 E = iota
	E2
	E3
)

func NewE() E {
	return E(0)
}

var __typeRefForE types.TypeRef

func (e E) TypeRef() types.TypeRef {
	return __typeRefForE
}

func init() {
	__typeRefForE = types.MakeTypeRef(__leafDepPackageInFile_leafDep_CachedRef, 1)
	types.RegisterEnum(__typeRefForE, builderForE, readerForE)
}

func builderForE(v uint32) types.Value {
	return E(v)
}

func readerForE(v types.Value) uint32 {
	return uint32(v.(E))
}

func (e E) Equals(other types.Value) bool {
	return e == other
}

func (e E) Ref() ref.Ref {
	throwaway := ref.Ref{}
	return types.EnsureRef(&throwaway, e)
}

func (e E) Chunks() []ref.Ref {
	return nil
}

func (e E) ChildValues() []types.Value {
	return nil
}
