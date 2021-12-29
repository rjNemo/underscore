package underscore

// Reduce combine a list of values into a single value.
// acc is the initial state, and each successive step of it should be returned by the reduction function.
func Reduce[T, P any](values []T, reduction func(T, P) P, acc P) P {
	for _, v := range values {
		acc = reduction(v, acc)
	}
	return acc
}
