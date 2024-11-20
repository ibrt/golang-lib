package memz

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

// PtrIfTrue returns a pointer to the given value of cond is true, nil otherwise.
func PtrIfTrue[T any](cond bool, v T) *T {
	if cond {
		return &v
	}
	return nil
}

// PtrZeroToNil returns a pointer to the given value if different from the zero-value, nil otherwise.
func PtrZeroToNil[T comparable](v T) *T {
	var z T
	if v == z {
		return nil
	}

	return &v
}

// PtrZeroToNilIfTrue returns a pointer to the given value if different from the zero-value and cond is true, nil otherwise.
func PtrZeroToNilIfTrue[T comparable](cond bool, v T) *T {
	var z T
	if !cond || v == z {
		return nil
	}

	return &v
}

// ValNilToZero returns the value of the given pointer if found, a zero-value of the same type otherwise.
func ValNilToZero[T any](v *T) T {
	if v == nil {
		var z T
		return z
	}

	return *v
}

// ValNilToDef returns the value of the given pointer if found, the given default value otherwise.
func ValNilToDef[T any](v *T, d T) T {
	if v == nil {
		return d
	}

	return *v
}
