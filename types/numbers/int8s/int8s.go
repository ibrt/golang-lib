package int8s

import (
	"sort"
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 8
)

// Ptr returns a pointer to the value.
func Ptr(v int8) *int8 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v int8) *int8 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v int8, def int8) *int8 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *int8, def int8) int8 {
	if v == nil {
		return def
	}
	return *v
}

// ParseDec parses a string as base 10 int8.
func Parse(v string) (int8, error) {
	p, err := strconv.ParseInt(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (int8)(p), nil
}

// Slice is a slice of values.
type Slice []int8

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
func SliceToMap(s []int8) map[int8]struct{} {
	m := make(map[int8]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[int8]struct{}) []int8 {
	s := make([]int8, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}
