package underscore

// Every returns true if all the values in the slice pass the predicate truth test.
// Short-circuits and stops traversing the slice if a false element is found.
func Every[T any](values []T, predicate func(T) bool) bool {
	for _, v := range values {
		if !predicate(v) {
			return false
		}
	}
	return true
}
