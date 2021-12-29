package underscore

// Map produces a new slice of values by mapping each value in the slice through
// a transform function.
func Map[T, P any](values []T, transform func(T) P) []P {
	res := make([]P, 0, len(values))
	for _, v := range values {
		res = append(res, transform(v))
	}
	return res
}
