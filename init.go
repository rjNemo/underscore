package underscore

// Init returns all elements except the last one, and the last element separately.
// Returns an empty slice and zero value if the input slice is empty.
// Also known as "uncons from the right" or "snoc" inverse.
func Init[T any](values []T) ([]T, T) {
	var last T
	if len(values) == 0 {
		return []T{}, last
	}
	if len(values) == 1 {
		return []T{}, values[0]
	}

	res := make([]T, len(values)-1)
	copy(res, values[:len(values)-1])
	return res, values[len(values)-1]
}
