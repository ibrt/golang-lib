package float64s

import (
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 64
)

// Ptr returns a pointer to the value.
func Ptr(v float64) *float64 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v float64) *float64 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v float64, def float64) *float64 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *float64, def float64) float64 {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 float64.
func Parse(v string) (float64, error) {
	p, err := strconv.ParseFloat(v, BitSize)
	if err != nil {
		return 0, err
	}
	return (float64)(p), nil
}
