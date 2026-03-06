package generic

import "slices"

// RemoveDuplicates removes duplicate elements from a slice of comparable values.
// It modifies the slice in place and returns the deduplicated slice with clipped capacity.
func RemoveDuplicates[S ~[]E, E comparable](s S) S {
	return RemoveDuplicatesFunc(s, func(a, b E) bool { return a == b })
}

// RemoveDuplicatesFunc is like [RemoveDuplicates] but uses an
// equality function to compare elements.
func RemoveDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) S {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); {
			if equal(s[i], s[j]) {
				s = slices.Delete(s, j, j+1)
				// Don't increment j; check the same position again after deletion
			} else {
				j++
			}
		}
	}

	return slices.Clip(s)
}

// ContainsDuplicates reports whether the slice contains duplicate elements.
// Equality is determined using the == operator.
func ContainsDuplicates[S ~[]E, E comparable](s S) bool {
	return ContainsDuplicatesFunc(s, func(a, b E) bool { return a == b })
}

// ContainsDuplicatesFunc is like [ContainsDuplicates] but uses an
// equality function to compare elements.
func ContainsDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) bool {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if equal(s[i], s[j]) {
				return true
			}
		}
	}

	return false
}

// GetDuplicateIndexes returns duplicate indexes for each element in the slice.
func GetDuplicateIndexes[S ~[]E, E comparable](s S) map[int][]int {
	return GetDuplicateIndexesFunc(s, func(a, b E) bool { return a == b })
}

// GetDuplicateIndexesFunc is like [GetDuplicateIndexes] but uses an
// equality function to compare elements.
func GetDuplicateIndexesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) map[int][]int {
	duplicates := make(map[int][]int)

	for i := range s {
		for j := range s {
			if i == j {
				continue
			}
			if equal(s[i], s[j]) {
				duplicates[i] = append(duplicates[i], j)
			}
		}
	}

	return duplicates
}

// RemoveDuplicatesByComparableKey removes duplicate elements from a slice using a derived comparable key.
// The key function extracts a key from each element; elements with repeated keys are skipped.
func RemoveDuplicatesByComparableKey[T any, K comparable](slice []T, fn func(T) K) []T {
	keys := make(map[K]struct{})
	result := []T{}

	for _, item := range slice {
		key := fn(item)
		if _, exists := keys[key]; !exists {
			keys[key] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
