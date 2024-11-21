package tplz

import (
	"bytes"
	"go/format"
	ttpl "text/template"

	"github.com/ibrt/golang-lib/errorz"
)

// ExecuteGolang executes a text template, formatting the result as Go code.
func ExecuteGolang(template *ttpl.Template, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, errorz.Wrap(err)
	}

	return format.Source(buf.Bytes())
}

// MustExecuteGolang is like ExecuteGolang but panics on error.
func MustExecuteGolang(template *ttpl.Template, data any) []byte {
	buf, err := ExecuteGolang(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}

// ParseAndExecuteGolang parses and executes a text template, formatting the result as Go code.
func ParseAndExecuteGolang(template string, data any) ([]byte, error) {
	parsedTemplate, err := ttpl.New("").Parse(template)
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	return ExecuteGolang(parsedTemplate, data)
}

// MustParseAndExecuteGolang is like ParseAndExecuteGolang but panics on error.
func MustParseAndExecuteGolang(template string, data any) []byte {
	buf, err := ParseAndExecuteGolang(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}
