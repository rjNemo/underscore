package underscore

import "cmp"

// Min returns the minimum value in the slice.
// Panics if values is empty.
// This function can currently only compare numbers reliably.
// This function uses operator <.
func Min[T cmp.Ordered](values []T) T {
	if len(values) == 0 {
		panic("underscore.Min: empty slice")
	}
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
