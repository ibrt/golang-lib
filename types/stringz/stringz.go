package stringz

import (
	"fmt"
	"sort"
)

// Ptr returns a pointer to the value.
func Ptr(v string) *string {
	return &v
}

// PtrEmptyToNil returns a pointer to the value, or nil if "".
func PtrEmptyToNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v string, def string) *string {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to "" if nil.
func Val(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *string, def string) string {
	if v == nil {
		return def
	}
	return *v
}

// Slice is a slice of values.
type Slice []string

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
func SliceToMap(s []string) map[string]struct{} {
	m := make(map[string]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[string]struct{}) []string {
	s := make([]string, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}

// SwapMap returns a copy of the map with keys and values swapped.
// Fails in case of duplicate values.
func SwapMap(m map[string]string) (map[string]string, error) {
	i := make(map[string]string, len(m))

	for k, v := range m {
		if _, ok := i[v]; ok {
			return nil, fmt.Errorf("duplicate value: %v", v)
		}
		i[v] = k
	}

	return i, nil
}

// SafeIndex returns "s[i]" if possible, an empty string otherwise.
func SafeIndex(s []string, i int) string {
	if s == nil || i < 0 || i >= len(s) {
		return ""
	}
	return s[i]
}

// SafeIndexPtr returns "s[i]" if possible, an nil otherwise.
func SafeIndexPtr(s []string, i int) *string {
	if s == nil || i < 0 || i >= len(s) {
		return nil
	}
	return Ptr(s[i])
}
