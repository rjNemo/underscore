package underscore

import "errors"

// ErrEmptySlice is returned when trying to get the first element of an empty slice
var ErrEmptySlice = errors.New("underscore: empty slice")

// First returns the first element of the slice.
// Returns an error if the slice is empty.
func First[T any](values []T) (T, error) {
	var zero T
	if len(values) == 0 {
		return zero, ErrEmptySlice
	}
	return values[0], nil
}

// FirstN returns the first n elements of the slice.
// If n is greater than the slice length, returns the entire slice.
// If n is less than or equal to 0, returns an empty slice.
func FirstN[T any](values []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(values) {
		res := make([]T, len(values))
		copy(res, values)
		return res
	}
	res := make([]T, n)
	copy(res, values[:n])
	return res
}
