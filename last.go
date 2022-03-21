package underscore

// Last returns the last element of the slice
func Last[T any](values []T) T {
	n := len(values)
	return values[n-1]
}
