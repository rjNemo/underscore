package underscore

// TakeWhile returns elements from the beginning of the slice while the predicate returns true.
// It stops at the first element where the predicate returns false.
func TakeWhile[T any](values []T, predicate func(T) bool) []T {
	for i, v := range values {
		if !predicate(v) {
			res := make([]T, i)
			copy(res, values[:i])
			return res
		}
	}
	// All elements satisfy predicate
	res := make([]T, len(values))
	copy(res, values)
	return res
}
