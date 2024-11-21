package tplz

import (
	"bytes"
	ttpl "text/template"

	"github.com/ibrt/golang-lib/errorz"
)

// ParseAndExecuteText parses and executes a text template.
func ParseAndExecuteText(template string, data any) ([]byte, error) {
	parsedTemplate, err := ttpl.New("").Parse(template)
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	buf := &bytes.Buffer{}
	if err := parsedTemplate.Execute(buf, data); err != nil {
		return nil, errorz.Wrap(err)
	}

	return buf.Bytes(), nil
}

// MustParseAndExecuteText is like ParseAndExecuteText but panics on error.
func MustParseAndExecuteText(template string, data any) []byte {
	buf, err := ParseAndExecuteText(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}

// ExecuteText executes a text template.
func ExecuteText(template *ttpl.Template, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, errorz.Wrap(err)
	}

	return buf.Bytes(), nil
}

// MustExecuteText is like ExecuteText but panics on error.
func MustExecuteText(template *ttpl.Template, data any) []byte {
	buf, err := ExecuteText(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}
