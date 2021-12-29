package underscore

// Filter looks through each value in the slice, returning a slice of all the values that pass a truth test (predicate).
func Filter[T any](values []T, predicate func(T) bool) (res []T) {
	for _, v := range values {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}
