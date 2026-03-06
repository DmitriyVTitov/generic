package generic

// Val returns value of pointer or default value if pointer is 'nil'.
func Val[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}

	return *ptr
}

// Ptr returns a pointer to the any given value when '&' does not work.
// In Go 1.26+ use built-in 'new' for expressions instead.
func Ptr[T any](val T) *T {
	return &val
}
