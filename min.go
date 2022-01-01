package underscore

import "constraints"

// Min returns the minimum value in the slice.
// This function can currently only compare numbers reliably.
// This function uses operator <.
func Min[T constraints.Ordered](values []T) T {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
