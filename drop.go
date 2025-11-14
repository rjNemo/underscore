package underscore

// Drop returns a new slice with the first n elements removed.
// If n is greater than or equal to the slice length, returns an empty slice.
// If n is less than or equal to 0, returns the original slice.
func Drop[T any](values []T, n int) []T {
	if n <= 0 {
		return values
	}
	if n >= len(values) {
		return []T{}
	}
	res := make([]T, len(values)-n)
	copy(res, values[n:])
	return res
}
