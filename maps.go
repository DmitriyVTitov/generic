package generic

// MapItem returns the value for the specified key from the map if it exists
// and the associated value can be asserted to type T. Otherwise, it returns
// the zero value of type T.
func MapItem[T any, K comparable, V any](m map[K]V, key K) T {
	var zero T

	if m == nil {
		return zero
	}

	if val, ok := m[key]; ok {
		if castVal, ok := any(val).(T); ok {
			return castVal
		}
	}

	return zero
}
