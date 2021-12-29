package underscore

func Each[T any](values []T, predicate func(T)) {
	for _, v := range values {
		predicate(v)
	}
}
