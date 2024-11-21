package underscore

// Unique returns a slice of unique values from the given slice.
func Unique[T comparable](values []T) (uniques []T) {
	seen := make(map[T]bool, 0)
	for _, v := range values {
		if _, ok := seen[v]; !ok {
			uniques = append(uniques, v)
			seen[v] = true
		}
	}
	return uniques
}
