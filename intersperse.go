package underscore

// Intersperse inserts a separator between each element of the slice.
// Returns an empty slice if the input is empty.
// Returns the original element if the input has only one element.
//
// Example: Intersperse([]int{1,2,3}, 0) → [1, 0, 2, 0, 3]
func Intersperse[T any](values []T, separator T) []T {
	if len(values) == 0 {
		return []T{}
	}
	if len(values) == 1 {
		return []T{values[0]}
	}

	// Result will have len(values) + (len(values)-1) elements
	res := make([]T, 0, len(values)*2-1)
	res = append(res, values[0])
	for i := 1; i < len(values); i++ {
		res = append(res, separator, values[i])
	}
	return res
}
