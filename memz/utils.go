package memz

import (
	"cmp"
	"fmt"
)

// Min returns the lowest between v1 and v2.
func Min[T cmp.Ordered](v1, v2 T) T {
	if cmp.Less(v1, v2) {
		return v1
	}

	return v2
}

// Max returns the highest between v1 and v2.
func Max[T cmp.Ordered](v1, v2 T) T {
	if cmp.Less(v1, v2) {
		return v2
	}

	return v1
}

// PredicateIsZeroValue returns true if v is the zero-value of its type.
func PredicateIsZeroValue[T comparable](v T) bool {
	var z T
	return v == z
}

// TransformSprintf stringifies values using fmt.Sprintf("%v").
func TransformSprintf[V any](v V) string {
	return fmt.Sprintf("%v", v)
}
