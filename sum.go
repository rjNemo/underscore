package underscore

import "golang.org/x/exp/constraints"

// Sum adds elements of the slice.
func Sum[T constraints.Ordered](values []T) (sum T) {
	for _, v := range values {
		sum += v
	}
	return sum
}
