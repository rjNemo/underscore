package underscore

import "slices"

// Any returns true if any of the values in the slice pass the predicate truth test.
// Short-circuits and stops traversing the slice if a true element is found.
func Any[T any](values []T, predicate func(T) bool) bool {
	return slices.ContainsFunc(values, predicate)
}
