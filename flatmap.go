package underscore

// Flatmap flatten the input slice element into the new slice. FlatMap maps every element with the help of a mapper function, then flattens the input slice element into the new slice.
func Flatmap[T any](values []T, mapper func(n T) []T) []T {
	res := make([]T, 0)
	for _, v := range values {
		vs := mapper(v)
		res = append(res, vs...)
	}
	return res
}
