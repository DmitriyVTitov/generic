package generic

// Val returns value of pointer or default value if pointer is 'nil'.
func Val[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}

	return *ptr
}

// Ptr returns a pointer to the given value.
// In Go 1.26+ use built-in 'new' instead.
func Ptr[T any](val T) *T {
	return &val
}
