package underscore

// Partition splits the slice into two slices: one whose elements all satisfy predicate
// and one whose elements all do not satisfy predicate.
func Partition[T any](values []T, predicate func(T) bool) ([]T, []T) {
	keep := make([]T, 0)
	reject := make([]T, 0)

	for _, v := range values {
		if predicate(v) {
			keep = append(keep, v)
		} else {
			reject = append(reject, v)
		}
	}
	return keep, reject
}
