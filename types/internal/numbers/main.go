package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ibrt/golang-lib/errors"
)

var tpl = template.Must(template.New("").Parse(`package {{ .Pkg }}

import (
	"sort"
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = {{ if eq .Size -1 }}32 << (^uint(0) >> 32 & 1){{ else }}{{ .Size }}{{ end }}
)

// Ptr returns a pointer to the value.
func Ptr(v {{ .Type }}) *{{ .Type }} {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v {{ .Type }}) *{{ .Type }} {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v {{ .Type }}, def {{ .Type }}) *{{ .Type }} {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *{{ .Type }}) {{ .Type }} {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *{{ .Type }}, def {{ .Type }}) {{ .Type }} {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 {{ .Type }}.
func Parse(v string) ({{ .Type }}, error) {
	p, err := {{ .ParseCall }}
	if err != nil {
		return 0, err
	}
	return ({{ .Type }})(p), nil
}

// Slice is a slice of values.
type Slice []{{ .Type }}

// Len implements the sort.Interface interface.
func (s Slice) Len() int {
	return len(s)
}

// Less implements the sort.Interface interface.
func (s Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap implements the sort.Interface interface.
func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort the slice.
func (s Slice) Sort() {
	sort.Sort(s)
}

// IsSorted returns true if the slice is sorted.
func (s Slice) IsSorted() bool {
	return sort.IsSorted(s)
}

// SliceToMap converts a slice to map.
func SliceToMap(s []{{ .Type }}) map[{{ .Type }}]struct{} {
	m := make(map[{{ .Type }}]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[{{ .Type }}]struct{}) []{{ .Type }} {
	s := make([]{{ .Type }}, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}
`))

var testTpl = template.Must(template.New("").Parse(`package {{ .Pkg }}_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/{{ .Pkg }}"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := {{ .Pkg }}.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, {{ .Type }}(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := {{ .Pkg }}.PtrZeroToNil(0)
	require.Nil(t, p)
	p = {{ .Pkg }}.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, {{ .Type }}(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := {{ .Pkg }}.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = {{ .Pkg }}.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, {{ .Type }}(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, {{ .Type }}(0), {{ .Pkg }}.Val(nil))
	require.Equal(t, {{ .Type }}(0), {{ .Pkg }}.Val({{ .Pkg }}.Ptr(0)))
	require.Equal(t, {{ .Type }}(1), {{ .Pkg }}.Val({{ .Pkg }}.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, {{ .Type }}(1), {{ .Pkg }}.ValDef(nil, 1))
	require.Equal(t, {{ .Type }}(0), {{ .Pkg }}.ValDef({{ .Pkg }}.Ptr(0), 1))
	require.Equal(t, {{ .Type }}(1), {{ .Pkg }}.ValDef({{ .Pkg }}.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := {{ .Pkg }}.Parse("10")
	require.NoError(t, err)
	require.Equal(t, {{ .Type }}(10), v)
	_, err = {{ .Pkg }}.Parse("")
	require.Error(t, err)
	_, err = {{ .Pkg }}.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []{{ .Type }}{2, 0, 3, 1, 4}
	require.False(t, {{ .Pkg }}.Slice(s).IsSorted())
	{{ .Pkg }}.Slice(s).Sort()
	require.Equal(t, []{{ .Type }}{0, 1, 2, 3, 4}, s)
	require.True(t, {{ .Pkg }}.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[{{ .Type }}]struct{}{}, {{ .Pkg }}.SliceToMap(nil))
	require.Equal(t, map[{{ .Type }}]struct{}{}, {{ .Pkg }}.SliceToMap([]{{ .Type }}{}))
	require.Equal(t, map[{{ .Type }}]struct{}{1: {}}, {{ .Pkg }}.SliceToMap([]{{ .Type }}{1}))
	require.Equal(t, map[{{ .Type }}]struct{}{1: {}, 2: {}}, {{ .Pkg }}.SliceToMap([]{{ .Type }}{1, 2}))
	require.Equal(t, map[{{ .Type }}]struct{}{1: {}, 2: {}}, {{ .Pkg }}.SliceToMap([]{{ .Type }}{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []{{ .Type }}{}, {{ .Pkg }}.MapToSlice(nil))
	require.Equal(t, []{{ .Type }}{}, {{ .Pkg }}.MapToSlice(map[{{ .Type }}]struct{}{}))
	require.Equal(t, []{{ .Type }}{1}, {{ .Pkg }}.MapToSlice(map[{{ .Type }}]struct{}{1: {}}))
	require.Equal(t, map[{{ .Type }}]struct{}{1: {}, 2: {}}, {{ .Pkg }}.SliceToMap({{ .Pkg }}.MapToSlice(map[{{ .Type }}]struct{}{1: {}, 2: {}})))
}
`))

type params struct {
	Type string
	Size int
}

func (p *params) Pkg() string {
	return p.Type + "s"
}

func (p *params) ParseCall() string {
	switch p.Type[0] {
	case 'i':
		return "strconv.ParseInt(v, 10, BitSize)"
	case 'u':
		return "strconv.ParseUint(v, 10, BitSize)"
	case 'f':
		return "strconv.ParseFloat(v, BitSize)"
	default:
		panic("unknown type")
	}
}

var genParams = []*params{
	{"int8", 8},
	{"int16", 16},
	{"int32", 32},
	{"int64", 64},
	{"int", -1},
	{"uint8", 8},
	{"uint16", 16},
	{"uint32", 32},
	{"uint64", 64},
	{"uint", -1},
	{"float32", 32},
	{"float64", 64},
}

func main() {
	fmt.Println("numbers: cleaning up...")

	errors.MaybeMustWrap(os.RemoveAll(filepath.Join("..", "numbers")))
	errors.MaybeMustWrap(os.MkdirAll(filepath.Join("..", "numbers"), 0777))

	fmt.Println("numbers: generating...")

	for _, params := range genParams {
		buf := &bytes.Buffer{}
		errors.MaybeMustWrap(tpl.Execute(buf, params))
		errors.MaybeMustWrap(os.MkdirAll(filepath.Join("..", "numbers", params.Pkg()), 0777))
		errors.MaybeMustWrap(ioutil.WriteFile(filepath.Join("..", "numbers", params.Pkg(), params.Pkg()+".go"), buf.Bytes(), 0666))

		buf = &bytes.Buffer{}
		errors.MaybeMustWrap(testTpl.Execute(buf, params))
		errors.MaybeMustWrap(ioutil.WriteFile(filepath.Join("..", "numbers", params.Pkg(), params.Pkg()+"_test.go"), buf.Bytes(), 0666))
	}
}
