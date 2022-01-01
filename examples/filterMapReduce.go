package examples

import (
	"fmt"

	u "github.com/rjNemo/underscore"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// filter even numbers from the slice
	evens := u.Filter(numbers, func(n int) bool { return n%2 == 0 })
	// square every number in the slice
	squares := u.Map(evens, func(n int) int { return n * n })
	// reduce to the sum
	res := u.Reduce(squares, func(n, acc int) int { return n + acc }, 0)

	fmt.Println(res)
}
