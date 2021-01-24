package float64s

import (
	"sort"
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

// Slice is a slice of values.
type Slice []float64

// Len implements the sort.Interface interface.
func (s Slice) Len() int {
	return len(s)
}

// Less implements the sort.Interface interface.
func (s Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap implements the sort.Interface interface.
func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort the slice.
func (s Slice) Sort() {
	sort.Sort(s)
}

// IsSorted returns true if the slice is sorted.
func (s Slice) IsSorted() bool {
	return sort.IsSorted(s)
}

// SliceToMap converts a slice to map.
func SliceToMap(s []float64) map[float64]struct{} {
	m := make(map[float64]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[float64]struct{}) []float64 {
	s := make([]float64, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}
