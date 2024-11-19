package jsonz

import (
	"bytes"
	"encoding/json"

	"github.com/ibrt/golang-lib/errorz"
)

// MustMarshal is like json.Marshal but panics on error.
func MustMarshal(v any) []byte {
	buf, err := json.Marshal(v)
	errorz.MaybeMustWrap(err)
	return buf
}

// MustMarshalString is like MustMarshal but returns a string.
func MustMarshalString(v any) string {
	return string(MustMarshal(v))
}

// MustMarshalIndent is like json.MarshalIndent but panics on error.
func MustMarshalIndent(v any, prefix, indent string) []byte {
	buf, err := json.MarshalIndent(v, prefix, indent)
	errorz.MaybeMustWrap(err)
	return buf
}

// MustMarshalIndentString is like MustMarshalIndent but returns a string.
func MustMarshalIndentString(v any, prefix, indent string) string {
	return string(MustMarshalIndent(v, prefix, indent))
}

// MustMarshalIndentDefault is like MustMarshalIndent with prefix = "" and indent = "  ".
func MustMarshalIndentDefault(v any) []byte {
	return MustMarshalIndent(v, "", "  ")
}

// MustMarshalIndentDefaultString is like MustMarshalIndentDefault with prefix = "" and indent = "  ".
func MustMarshalIndentDefaultString(v any) string {
	return MustMarshalIndentString(v, "", "  ")
}

// MustIndent is like json.Indent but panics on error.
func MustIndent(src []byte, prefix, indent string) []byte {
	dst := &bytes.Buffer{}
	errorz.MaybeMustWrap(json.Indent(dst, src, prefix, indent))
	return dst.Bytes()
}

// MustIndentString is like MustIndent but returns a string.
func MustIndentString(src []byte, prefix, indent string) string {
	return string(MustIndent(src, prefix, indent))
}

// MustIndentDefault is like MustIndent with prefix = "" and indent = "  ".
func MustIndentDefault(src []byte) []byte {
	return MustIndent(src, "", "  ")
}

// MustIndentDefaultString is like MustIndentString with prefix = "" and indent = "  ".
func MustIndentDefaultString(src []byte) string {
	return MustIndentString(src, "", "  ")
}

// Unmarshal is like json.Unmarshal but instantiates the target using a generic type.
func Unmarshal[T any](data []byte) (*T, error) {
	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, errorz.Wrap(err)
	}

	return &t, nil
}

// MustUnmarshal is like Unmarshal but panics on error.
func MustUnmarshal[T any](data []byte) *T {
	t, err := Unmarshal[T](data)
	errorz.MaybeMustWrap(err)
	return t
}

// UnmarshalString is like Unmarshal but accepts a string.
func UnmarshalString[T any](data string) (*T, error) {
	var t T
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return nil, errorz.Wrap(err)
	}

	return &t, nil
}

// MustUnmarshalString is like UnmarshalString but panics on error.
func MustUnmarshalString[T any](data string) *T {
	t, err := UnmarshalString[T](data)
	errorz.MaybeMustWrap(err)
	return t
}
