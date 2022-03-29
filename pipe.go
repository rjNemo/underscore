package underscore

import (
	"golang.org/x/exp/constraints"
)

type Pipe[T constraints.Ordered] struct {
	Value []T
}

// NewPipe starts a Pipe. All future method calls will return Pipe structs. When you've finished the computation,
// call Value to retrieve the final value.
//
// Methods not returning a slice such as Reduce, All, Any, will break the Pipe and return Value instantly.
func NewPipe[T constraints.Ordered](value []T) Pipe[T] {
	return Pipe[T]{Value: value}
}

// All returns true if all the values in the slice pass the predicate truth test.
// Short-circuits and stops traversing the slice if a false element is found.
// Breaks the Pipe.
func (c Pipe[T]) All(predicate func(T) bool) bool {
	return All(c.Value, predicate)
}

// Any returns true if any of the values in the slice pass the predicate truth test.
// Short-circuits and stops traversing the slice if a true element is found.
// Breaks the Pipe.
func (c Pipe[T]) Any(predicate func(T) bool) bool {
	return Any(c.Value, predicate)
}

// Contains returns true if the value is present in the slice and breaks the Pipe.
func (c Pipe[T]) Contains(value T) bool {
	return Contains(c.Value, value)
}

// Each iterates over a slice of elements, yielding each in turn to an action function.
// Breaks the Pipe.
func (c Pipe[T]) Each(action func(T)) {
	Each(c.Value, action)
}

// Filter looks through each value in the slice, returning a slice of all the values that pass a truth test (predicate).
func (c Pipe[T]) Filter(predicate func(n T) bool) Pipe[T] {
	return Pipe[T]{Value: Filter(c.Value, predicate)}
}

// Find looks through each value in the slice, returning the first one that passes a truth test (predicate),
// or the default value for the type and an error if no value passes the test.
// The function returns as soon as it finds an acceptable element, and doesn't traverse the entire slice.
// Breaks the Pipe.
func (c Pipe[T]) Find(predicate func(n T) bool) (T, error) {
	return Find(c.Value, predicate)
}

// Map produces a new slice of values by mapping each value in the slice through
// a transform function.
//
// TODO: Move from T to P.
func (c Pipe[T]) Map(transform func(n T) T) Pipe[T] {
	return Pipe[T]{Value: Map(c.Value, transform)}
}

// Max returns the maximum value in the slice.
// This function can currently only compare numbers reliably.
// This function uses operator <.
// Breaks the Pipe.
func (c Pipe[T]) Max() T {
	return Max(c.Value)
}

// Min returns the minimum value in the slice.
// This function can currently only compare numbers reliably.
// This function uses operator <.
// Breaks the Pipe.
func (c Pipe[T]) Min() T {
	return Min(c.Value)
}

// Partition splits the slice into two slices: one whose elements all satisfy predicate
// and one whose elements all do not satisfy predicate.
// Breaks the Pipe.
func (c Pipe[T]) Partition(predicate func(T) bool) ([]T, []T) {
	return Partition(c.Value, predicate)
}

// Reduce combine a list of values into a single value and breaks the Pipe.
// acc is the initial state, and each successive step of it should be returned by the reduction function.
func (c Pipe[T]) Reduce(reducer func(n, acc T) T, acc T) T {
	return Reduce(c.Value, reducer, acc)
}
