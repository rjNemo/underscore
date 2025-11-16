package underscore

// Replicate creates a slice containing count copies of value.
// Returns an empty slice if count is less than or equal to 0.
//
// Example: Replicate(3, "hello") → ["hello", "hello", "hello"]
func Replicate[T any](count int, value T) []T {
	if count <= 0 {
		return []T{}
	}

	res := make([]T, count)
	for i := range res {
		res[i] = value
	}
	return res
}
