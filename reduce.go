package underscore

func Reduce[T, P any](values []T, predicate func(T, P) P, acc P) P {
	for _, v := range values {
		acc = predicate(v, acc)
	}
	return acc
}
