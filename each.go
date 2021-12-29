package underscore

// Each iterates over a slice of elements, yielding each in turn to an action function.
func Each[T any](values []T, action func(T)) {
	for _, v := range values {
		action(v)
	}
}
