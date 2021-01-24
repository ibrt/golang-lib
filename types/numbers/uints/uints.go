package uints

import (
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 32 << (^uint(0) >> 32 & 1)
)

// Ptr returns a pointer to the value.
func Ptr(v uint) *uint {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v uint) *uint {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v uint, def uint) *uint {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *uint, def uint) uint {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 uint.
func Parse(v string) (uint, error) {
	p, err := strconv.ParseUint(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (uint)(p), nil
}
