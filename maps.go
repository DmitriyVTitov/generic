package generic

// MapItem returns the value of the specified key from the map if it exists
// and is of type T. Otherwise, it returns the zero value of type T.
func MapItem[T any, Key comparable, Value any](m map[Key]Value, key Key) T {
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
