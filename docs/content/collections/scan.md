---
title: "Scan"
date: 2025-01-16T00:00:00-00:00
---

`Scan` is like Reduce but returns all intermediate accumulator values. Also known as prefix scan or cumulative fold. Useful for tracking running totals, running maximums, or other cumulative operations.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	// Running sum
	nums := []int{1, 2, 3, 4}
	add := func(acc, n int) int { return acc + n }
	fmt.Println(u.Scan(nums, 0, add)) // [1, 3, 6, 10]

	// Running maximum
	values := []int{3, 1, 4, 1, 5, 9, 2}
	max := func(acc, n int) int {
		if n > acc {
			return n
		}
		return acc
	}
	fmt.Println(u.Scan(values, 0, max)) // [3, 3, 4, 4, 5, 9, 9]

	// String concatenation
	words := []string{"hello", "world", "!"}
	concat := func(acc, s string) string { return acc + s }
	fmt.Println(u.Scan(words, "", concat)) // ["hello", "helloworld", "helloworld!"]
}
```
