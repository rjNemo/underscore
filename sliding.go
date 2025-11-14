package underscore

// Sliding creates a sliding window view of the slice with the specified window size.
// Returns an empty slice if size is less than or equal to 0.
// Returns an empty slice if size is greater than the slice length.
//
// Example: Sliding([]int{1,2,3,4,5}, 3) → [[1,2,3], [2,3,4], [3,4,5]]
func Sliding[T any](values []T, size int) [][]T {
	if size <= 0 || size > len(values) {
		return [][]T{}
	}

	windowCount := len(values) - size + 1
	res := make([][]T, 0, windowCount)

	for i := 0; i <= len(values)-size; i++ {
		window := make([]T, size)
		copy(window, values[i:i+size])
		res = append(res, window)
	}

	return res
}
