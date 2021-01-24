package float32s

import (
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 32
)

// Ptr returns a pointer to the value.
func Ptr(v float32) *float32 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v float32) *float32 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v float32, def float32) *float32 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *float32, def float32) float32 {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 float32.
func Parse(v string) (float32, error) {
	p, err := strconv.ParseFloat(v, BitSize)
	if err != nil {
		return 0, err
	}
	return (float32)(p), nil
}
