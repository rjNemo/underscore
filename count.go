package underscore

// Count returns the number of elements in the slice that satisfy the predicate.
// example: Count([]int{1,2,3,4,5}, func(n int) bool { return n%2 == 0 }) // 2
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, item := range slice {
		if predicate(item) {
			count++
		}
	}
	return count
}
