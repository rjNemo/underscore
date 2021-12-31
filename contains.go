package underscore

// Contains returns true if the value is present in the slice
func Contains[T numbers](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
