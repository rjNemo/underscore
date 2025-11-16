package underscore

// DropWhile drops elements from the beginning of the slice while the predicate returns true.
// It returns the remaining elements starting from the first element where the predicate returns false.
func DropWhile[T any](values []T, predicate func(T) bool) []T {
	for i, v := range values {
		if !predicate(v) {
			res := make([]T, len(values)-i)
			copy(res, values[i:])
			return res
		}
	}
	// All elements satisfy predicate, return empty slice
	return []T{}
}
