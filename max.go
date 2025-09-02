package underscore

import "cmp"

// Max returns the maximum value in the slice.
// This function can currently only compare numbers reliably.
// This function uses operator <.
func Max[T cmp.Ordered](values []T) T {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
