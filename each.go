package underscore

// Each iterates over a slice of elements, yielding each in turn to an action function.
// Returns the slice for chaining.
func Each[T any](values []T, action func(T)) []T {
	for _, v := range values {
		action(v)
	}
	return values
}
