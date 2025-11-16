---
title: "FoldRight"
date: 2025-01-16T00:00:00-00:00
---

`FoldRight` is like Reduce but processes elements from right to left. Also known as foldr in Haskell. Important for non-associative operations where the order of evaluation matters.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	// Subtraction is non-associative
	nums := []int{1, 2, 3}

	// FoldRight: 1 - (2 - (3 - 0)) = 1 - (2 - 3) = 1 - (-1) = 2
	result := u.FoldRight(nums, 0, func(n, acc int) int { return n - acc })
	fmt.Println(result) // 2

	// Compare with Reduce (left fold): (0 - 1) - 2 - 3 = -6
	leftResult := u.Reduce(nums, func(n, acc int) int { return acc - n }, 0)
	fmt.Println(leftResult) // -6

	// Building a list in order
	buildList := u.FoldRight(nums, []int{}, func(n int, acc []int) []int {
		return append([]int{n}, acc...)
	})
	fmt.Println(buildList) // [1, 2, 3]

	// String concatenation
	words := []string{"a", "b", "c"}
	concat := u.FoldRight(words, "", func(s, acc string) string { return s + acc })
	fmt.Println(concat) // "abc"
}
```
