package underscore

type Chain[T any] struct {
	Value []T
}

func NewChain[T any](value []T) Chain[T] {
	return Chain[T]{Value: value}
}

func (c Chain[T]) Filter(predicate func(n T) bool) Chain[T] {
	return Chain[T]{Value: Filter(c.Value, predicate)}
}

func (c Chain[T]) Map(transform func(n T) T) Chain[T] {
	return Chain[T]{Value: Map(c.Value, transform)}
}

func (c Chain[T]) Reduce(reducer func(n, acc T) T, initialValue T) T {
	return Reduce(c.Value, reducer, initialValue)
}
