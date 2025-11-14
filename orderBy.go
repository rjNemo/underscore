package underscore

import "slices"

// OrderBy orders a slice by a field value within a struct, the predicate allows you
// to pick the fields you want to orderBy. Use > for ASC or < for DESC
// Uses O(n log n) sorting algorithm. Mutates the input slice.
//
//	func (left Person, right Person) bool { return left.Age > right.Age }
func OrderBy[T any](list []T, predicate func(T, T) bool) []T {
	slices.SortFunc(list, func(a, b T) int {
		if predicate(a, b) {
			return 1
		}
		if predicate(b, a) {
			return -1
		}
		return 0
	})
	return list
}
