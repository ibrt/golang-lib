package float32s

import (
	"sort"
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

// Parse parses a string as base 10 float32.
func Parse(v string) (float32, error) {
	p, err := strconv.ParseFloat(v, BitSize)
	if err != nil {
		return 0, err
	}
	return (float32)(p), nil
}

// Slice is a slice of values.
type Slice []float32

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
func SliceToMap(s []float32) map[float32]struct{} {
	m := make(map[float32]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[float32]struct{}) []float32 {
	s := make([]float32, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}
