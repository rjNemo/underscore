package underscore

func Filter[T any](values []T, predicate func(T) bool) (res []T) {
	for _, v := range values {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}
