package examples

import (
	u "github.com/rjNemo/underscore"
)

func piping() int {
	return u.NewPipe([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).
		// filter even numbers from the slice
		Filter(func(n int) bool { return n%2 == 0 }).
		// square every number in the slice
		Map(func(n int) int { return n * n }).
		// reduce the slice to its sum
		Reduce(func(n, acc int) int { return n + acc }, 0)
}
