package underscore

// RemoveAt returns a new slice with the element at the given index removed.
// Returns original slice if index is out of bounds.
func RemoveAt[T any](values []T, index int) []T {
	if index < 0 || index >= len(values) {
		return values
	}
	res := make([]T, 0, len(values)-1)
	for i, value := range values {
		if i != index {
			res = append(res, value)
		}
	}
	return res
}
