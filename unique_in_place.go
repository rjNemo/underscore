package underscore

// UniqueInPlace removes duplicate elements from the slice in place, preserving order.
// It returns the shortened slice containing the first occurrence of each value.
func UniqueInPlace[T comparable](values []T) []T {
	seen := make(map[T]struct{}, len(values))
	w := 0
	for _, v := range values {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		values[w] = v
		w++
	}
	return values[:w]
}
