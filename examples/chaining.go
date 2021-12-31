package examples

import (
	"fmt"

	u "github.com/rjNemo/underscore"
)

func main() {
	sum := u.NewChain([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).
		// filter even numbers from the slice
		Filter(func(n int) bool { return n%2 == 0 }).
		// square every number in the slice
		Map(func(n int) int { return n * n }).
		// reduce to the sum
		Reduce(func(n, acc int) int { return n + acc }, 0)

	fmt.Println(sum)
}
