package underscore

// Last returns the last element of the slice.
// Panics if the slice is empty.
func Last[T any](values []T) T {
	if len(values) == 0 {
		panic("underscore.Last: empty slice")
	}
	return values[len(values)-1]
}
