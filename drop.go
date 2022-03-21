package underscore

func Drop[T any](values []T, index int) (rest []T) {
	for i, value := range values {
		if i != index {
			rest = append(rest, value)
		}
	}
	return rest
}
