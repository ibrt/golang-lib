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
	"fmt"
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
func ParseDec(v string) ({{ .Type }}, error) {
	p, err := strconv.Parse{{ if .IsSigned }}Int{{ else }}Uint{{ end }}(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return ({{ .Type }})(p), nil
}

// ParseHex parses a string as base 16 {{ .Type }}.
func ParseHex(v string) ({{ .Type }}, error) {
	p, err := strconv.Parse{{ if .IsSigned }}Int{{ else }}Uint{{ end }}(v, 16, BitSize)
	if err != nil {
		return 0, err
	}
	return ({{ .Type }})(p), nil
}

// StrDec interprets the value as base 10 and converts it to string.
func StrDec(v {{ .Type }}) string {
	return fmt.Sprintf("%d", v)
}

// StrHex interprets the value as base 16 and converts it to string.
func StrHex(v {{ .Type }}) string {
	return fmt.Sprintf("%x", v)
}
`))

var testTpl = template.Must(template.New("").Parse(`package {{ .Pkg }}_test

import (
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/{{ .Pkg }}"
	"github.com/stretchr/testify/require"
)

const Max = {{ .Max }}

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

func TestParseDec(t *testing.T) {
	v, err := {{ .Pkg }}.ParseDec("10")
	require.NoError(t, err)
	require.Equal(t, {{ .Type }}(10), v)
	v, err = {{ .Pkg }}.ParseDec(fmt.Sprintf("%d", Max))
	require.NoError(t, err)
	require.Equal(t, {{ .Type }}(Max), v)
	_, err = {{ .Pkg }}.ParseDec("")
	require.Error(t, err)
	_, err = {{ .Pkg }}.ParseDec("A")
	require.Error(t, err)
}

func TestParseHex(t *testing.T) {
	v, err := {{ .Pkg }}.ParseHex("20")
	require.NoError(t, err)
	require.Equal(t, {{ .Type }}(0x20), v)
	v, err = {{ .Pkg }}.ParseHex(fmt.Sprintf("%x", Max))
	require.NoError(t, err)
	require.Equal(t, {{ .Type }}(Max), v)
	v, err = {{ .Pkg }}.ParseHex(fmt.Sprintf("%X", Max))
	require.NoError(t, err)
	require.Equal(t, {{ .Type }}(Max), v)
	_, err = {{ .Pkg }}.ParseHex("")
	require.Error(t, err)
}

func TestStrDec(t *testing.T) {
	require.Equal(t, "10", {{ .Pkg }}.StrDec(10))
}

func TestStrHex(t *testing.T) {
	require.Equal(t, "10", {{ .Pkg }}.StrHex(0x10))
	require.Equal(t, "a", {{ .Pkg }}.StrHex(0xA))
	require.Equal(t, "11", {{ .Pkg }}.StrHex(0x11))
}
`))

type params struct {
	Type     string
	IsSigned bool
	Size     int
}

func (p *params) Pkg() string {
	return p.Type + "s"
}

func (p *params) Max() string {
	if p.IsSigned {
		return fmt.Sprintf("1<<(%v.BitSize-1) - 1", p.Pkg())
	}
	return fmt.Sprintf("1<<%v.BitSize - 1", p.Pkg())
}

var genParams = []*params{
	{"int8", true, 8},
	{"int16", true, 16},
	{"int32", true, 32},
	{"int64", true, 64},
	{"int", true, -1},
	{"uint8", true, 8},
	{"uint16", true, 16},
	{"uint32", true, 32},
	{"uint64", true, 64},
	{"uint", true, -1},
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
