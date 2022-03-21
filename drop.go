package underscore

// Drop returns the rest of the elements in a slice.
// Pass an index to return the values of the slice from that index onward.
func Drop[T any](values []T, index int) (rest []T) {
	for i, value := range values {
		if i != index {
			rest = append(rest, value)
		}
	}
	return rest
}
