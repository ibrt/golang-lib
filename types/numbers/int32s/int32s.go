package int32s

import (
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 32
)

// Ptr returns a pointer to the value.
func Ptr(v int32) *int32 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v int32) *int32 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v int32, def int32) *int32 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *int32, def int32) int32 {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 int32.
func Parse(v string) (int32, error) {
	p, err := strconv.ParseInt(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (int32)(p), nil
}
