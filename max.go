package underscore

import "cmp"

// Max returns the maximum value in the slice.
// Panics if values is empty.
// This function can currently only compare numbers reliably.
// This function uses operator <.
func Max[T cmp.Ordered](values []T) T {
	if len(values) == 0 {
		panic("underscore.Max: empty slice")
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
