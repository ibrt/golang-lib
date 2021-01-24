package ints

import (
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 32 << (^uint(0) >> 32 & 1)
)

// Ptr returns a pointer to the value.
func Ptr(v int) *int {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v int) *int {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v int, def int) *int {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *int, def int) int {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 int.
func Parse(v string) (int, error) {
	p, err := strconv.ParseInt(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (int)(p), nil
}
