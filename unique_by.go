package underscore

// UniqueBy returns a slice of unique values from the given slice using a key selector.
// The first occurrence of each key is kept and order is preserved.
func UniqueBy[T any, K comparable](values []T, key func(T) K) (uniques []T) {
	seen := make(map[K]struct{})
	for _, v := range values {
		k := key(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		uniques = append(uniques, v)
	}
	return uniques
}
