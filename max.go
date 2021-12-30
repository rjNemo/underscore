package underscore

func Max[T numbers](values []T) T {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
