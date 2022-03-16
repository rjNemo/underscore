package chain

import (
	"golang.org/x/exp/constraints"

	u "github.com/rjNemo/underscore"
)

type Chain[T constraints.Ordered] struct {
	Value []T
}

// Of starts a Chain. All future method calls will return Chain structs. When you've finished the computation,
// call Value to retrieve the final value.
//
// Methods not returning a slice such as Reduce, All, Any, will break the chain and return Value instantly.
func Of[T constraints.Ordered](value []T) Chain[T] {
	return Chain[T]{Value: value}
}

// All returns true if all the values in the slice pass the predicate truth test.
// Short-circuits and stops traversing the slice if a false element is found.
// Breaks the Chain.
func (c Chain[T]) All(predicate func(T) bool) bool {
	return u.All(c.Value, predicate)
}

// Any returns true if any of the values in the slice pass the predicate truth test.
// Short-circuits and stops traversing the slice if a true element is found.
// Breaks the Chain.
func (c Chain[T]) Any(predicate func(T) bool) bool {
	return u.Any(c.Value, predicate)
}

// Contains returns true if the value is present in the slice and breaks the Chain.
func (c Chain[T]) Contains(value T) bool {
	return u.Contains(c.Value, value)
}

// Each iterates over a slice of elements, yielding each in turn to an action function.
// Breaks the Chain.
func (c Chain[T]) Each(action func(T)) {
	u.Each(c.Value, action)
}

// Filter looks through each value in the slice, returning a slice of all the values that pass a truth test (predicate).
func (c Chain[T]) Filter(predicate func(n T) bool) Chain[T] {
	return Chain[T]{Value: u.Filter(c.Value, predicate)}
}

// Find looks through each value in the slice, returning the first one that passes a truth test (predicate),
// or the default value for the type and an error if no value passes the test.
// The function returns as soon as it finds an acceptable element, and doesn't traverse the entire slice.
// Breaks the Chain.
func (c Chain[T]) Find(predicate func(n T) bool) (T, error) {
	return u.Find(c.Value, predicate)
}

// Map produces a new slice of values by mapping each value in the slice through
// a transform function.
//
// TODO: Move from T to P.
func (c Chain[T]) Map(transform func(n T) T) Chain[T] {
	return Chain[T]{Value: u.Map(c.Value, transform)}
}

// Max returns the maximum value in the slice.
// This function can currently only compare numbers reliably.
// This function uses operator <.
// Breaks the Chain.
func (c Chain[T]) Max() T {
	return u.Max(c.Value)
}

// Min returns the minimum value in the slice.
// This function can currently only compare numbers reliably.
// This function uses operator <.
// Breaks the Chain.
func (c Chain[T]) Min() T {
	return u.Min(c.Value)
}

// Partition splits the slice into two slices: one whose elements all satisfy predicate
// and one whose elements all do not satisfy predicate.
// Breaks the Chain.
func (c Chain[T]) Partition(predicate func(T) bool) ([]T, []T) {
	return u.Partition(c.Value, predicate)
}

// Reduce combine a list of values into a single value and breaks the Chain.
// acc is the initial state, and each successive step of it should be returned by the reduction function.
func (c Chain[T]) Reduce(reducer func(n, acc T) T, acc T) T {
	return u.Reduce(c.Value, reducer, acc)
}
