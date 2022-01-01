package underscore

type Chain[T any] struct {
	Value []T
}

// NewChain starts a Chain. All future method calls will return Chain structs. When you've finished the computation,
// call Value to retrieve the final value.
//
// Methods not returning a slice such as Reduce, Every, Some, will break the chain and return Value instantly.
func NewChain[T any](value []T) Chain[T] {
	return Chain[T]{Value: value}
}

// Filter looks through each value in the slice, returning a slice of all the values that pass a truth test (predicate).
func (c Chain[T]) Filter(predicate func(n T) bool) Chain[T] {
	return Chain[T]{Value: Filter(c.Value, predicate)}
}

// Map produces a new slice of values by mapping each value in the slice through
// a transform function.
//
// TODO: Move from T to P.
func (c Chain[T]) Map(transform func(n T) T) Chain[T] {
	return Chain[T]{Value: Map(c.Value, transform)}
}

// Reduce combine a list of values into a single value and breaks the Chain.
// acc is the initial state, and each successive step of it should be returned by the reduction function.
func (c Chain[T]) Reduce(reducer func(n, acc T) T, acc T) T {
	return Reduce(c.Value, reducer, acc)
}
