package uint32s

import (
	"fmt"
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 32
)

// Ptr returns a pointer to the value.
func Ptr(v uint32) *uint32 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v uint32) *uint32 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v uint32, def uint32) *uint32 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *uint32, def uint32) uint32 {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 uint32.
func ParseDec(v string) (uint32, error) {
	p, err := strconv.ParseInt(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (uint32)(p), nil
}

// ParseHex parses a string as base 16 uint32.
func ParseHex(v string) (uint32, error) {
	p, err := strconv.ParseInt(v, 16, BitSize)
	if err != nil {
		return 0, err
	}
	return (uint32)(p), nil
}

// StrDec interprets the value as base 10 and converts it to string.
func StrDec(v uint32) string {
	return fmt.Sprintf("%d", v)
}

// StrHex interprets the value as base 16 and converts it to string.
func StrHex(v uint32) string {
	return fmt.Sprintf("%x", v)
}
