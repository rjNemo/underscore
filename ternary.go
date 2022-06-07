package underscore

func Ternary[T any](condition bool, pos, neg T) T {
	if condition {
		return pos
	}
	return neg
}
