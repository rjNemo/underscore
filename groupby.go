package underscore

// GroupBy splits a slice into a map[K][]V grouped by the result of the iterator function.
func GroupBy[K comparable, V any](values []V, f func(V) K) map[K][]V {
	res := make(map[K][]V, len(values)/10)
	for _, v := range values {
		k := f(v)
		if r, ok := res[k]; ok {
			res[k] = append(r, v)
		} else {
			res[k] = []V{v}
		}
	}
	return res
}
