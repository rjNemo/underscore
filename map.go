package underscore

func Map[T, P any](values []T, predicate func(T) P) []P {
	res := make([]P, 0, len(values))
	for _, v := range values {
		res = append(res, predicate(v))
	}
	return res
}
