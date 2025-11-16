package underscore

// Scan is like Reduce but returns all intermediate accumulator values.
// Also known as prefix scan or cumulative fold.
//
// Example: Scan([]int{1,2,3,4}, 0, func(acc, n int) int { return acc + n }) → [1, 3, 6, 10]
func Scan[T, P any](values []T, acc P, fn func(P, T) P) []P {
	if len(values) == 0 {
		return []P{}
	}

	res := make([]P, 0, len(values))
	for _, v := range values {
		acc = fn(acc, v)
		res = append(res, acc)
	}
	return res
}
