package uint8s

import (
	"fmt"
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 8
)

// Ptr returns a pointer to the value.
func Ptr(v uint8) *uint8 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v uint8) *uint8 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v uint8, def uint8) *uint8 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *uint8, def uint8) uint8 {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 uint8.
func ParseDec(v string) (uint8, error) {
	p, err := strconv.ParseInt(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (uint8)(p), nil
}

// ParseHex parses a string as base 16 uint8.
func ParseHex(v string) (uint8, error) {
	p, err := strconv.ParseInt(v, 16, BitSize)
	if err != nil {
		return 0, err
	}
	return (uint8)(p), nil
}

// StrDec interprets the value as base 10 and converts it to string.
func StrDec(v uint8) string {
	return fmt.Sprintf("%d", v)
}

// StrHex interprets the value as base 16 and converts it to string.
func StrHex(v uint8) string {
	return fmt.Sprintf("%x", v)
}
