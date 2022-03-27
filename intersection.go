package underscore

// Intersection computes the list of values that are the intersection of all the slices.
// Each value in the result is present in each of the slices.
func Intersection[T comparable](a, b []T) (res []T) {
	for _, n := range a {
		if Contains(b, n) {
			res = append(res, n)
		}
	}
	return res
}
