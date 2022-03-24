package underscore

// Difference Returns a copy of the array with all instances of the values that are not present in the other array.
func Difference[T comparable](slice, other []T) []T {
	return Filter(slice, func(n T) bool {
		return !Contains(other, n)
	})
}
