package underscore

// Some returns true if any of the values in the slice pass the predicate truth test.
// Short-circuits and stops traversing the slice if a true element is found.
func Some[T any](values []T, predicate func(T) bool) bool {
	for _, v := range values {
		if predicate(v) {
			return true
		}
	}
	return false
}
