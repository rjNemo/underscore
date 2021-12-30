package examples

import (
	"fmt"

	u "github.com/rjNemo/underscore"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// filter even numbers from the slice
	isEven := func(n int) bool { return n%2 == 0 }
	evens := u.Filter(numbers, isEven)
	// square every numbers numbers in the slice
	toSquare := func(n int) int { return n * n }
	squares := u.Map(evens, toSquare)
	// reduce to the sum
	sum := func(n, acc int) int { return n + acc }
	res := u.Reduce(squares, sum, 0)

	fmt.Println(res)
}
