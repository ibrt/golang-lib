package memz

// MergeMaps returns a new map built by setting all key/value pairs from the given maps in order.
func MergeMaps[K comparable, V any](mm ...map[K]V) map[K]V {
	length := 0

	for _, m := range mm {
		length += len(m)
	}

	out := make(map[K]V, length)

	for _, m := range mm {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

// ShallowCopyMap makes a shallow copy of a map.
func ShallowCopyMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V, len(m))

	for k, v := range m {
		out[k] = v
	}

	return out
}

// FilterMap copies a map including only (k, v) pairs for which the predicate returns true.
func FilterMap[K comparable, V any](m map[K]V, f func(k K, v V) bool) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V, len(m))

	for k, v := range m {
		if f(k, v) {
			out[k] = v
		}
	}

	return out
}
