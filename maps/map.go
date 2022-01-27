package maps

type M[K comparable, V any] map[K]V

// Map produces a new slice of values by mapping each value in the slice through
// a transform function.
func Map[K, Q comparable, V, W any](m M[K, V], f func(K, V) M[Q, W]) M[Q, W] {
	res := make(M[Q, W], len(m))
	for k, v := range m {
		mm := f(k, v)
		for k2, v2 := range mm {
			res[k2] = v2
		}
	}
	return res
}
