package underscore

import "slices"

// Contains returns true if the value is present in the slice
func Contains[T comparable](values []T, value T) bool {
	return slices.Contains(values, value)
}

// ContainsBy returns true if any element in the slice satisfies the predicate.
func ContainsBy[T any](values []T, predicate func(T) bool) bool {
	return slices.ContainsFunc(values, predicate)
}
