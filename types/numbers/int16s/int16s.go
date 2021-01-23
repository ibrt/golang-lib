package int16s

import (
	"fmt"
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 16
)

// Ptr returns a pointer to the value.
func Ptr(v int16) *int16 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v int16) *int16 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v int16, def int16) *int16 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *int16, def int16) int16 {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 int16.
func ParseDec(v string) (int16, error) {
	p, err := strconv.ParseInt(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (int16)(p), nil
}

// ParseHex parses a string as base 16 int16.
func ParseHex(v string) (int16, error) {
	p, err := strconv.ParseInt(v, 16, BitSize)
	if err != nil {
		return 0, err
	}
	return (int16)(p), nil
}

// StrDec interprets the value as base 10 and converts it to string.
func StrDec(v int16) string {
	return fmt.Sprintf("%d", v)
}

// StrHex interprets the value as base 16 and converts it to string.
func StrHex(v int16) string {
	return fmt.Sprintf("%x", v)
}
