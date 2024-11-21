package underscore

import "golang.org/x/exp/constraints"

// Sum adds elements of the slice.
func Sum[T constraints.Ordered](values []T) (sum T) {
	for _, v := range values {
		sum += v
	}
	return sum
}

// SumMap sums the values you select from your struct, basically a sort cut instead of
// having to perform a [Map] followed by a [Sum].
func SumMap[T any, R constraints.Ordered](list []T, selector func(T) R) (sum R) {
	for _, v := range list {
		sum += selector(v)
	}

	return sum
}
