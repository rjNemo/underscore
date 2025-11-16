package underscore

// Tap applies a function to each element for side effects (like debugging/logging)
// and returns the original slice unchanged. Useful for debugging pipelines.
//
// Example: Tap([]int{1,2,3}, func(n int) { fmt.Println(n) }) → [1,2,3] (and prints each)
func Tap[T any](values []T, fn func(T)) []T {
	for _, v := range values {
		fn(v)
	}
	return values
}
