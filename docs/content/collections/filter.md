---
title: "Filter"
date: 2022-03-21T13:31:21-04:00
---

`Filter` looks through each value in the slice, returning a slice of all the values that pass a truth test (predicate).

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(n int) bool { return n%2 == 0 }
	fmt.Println(u.Filter(nums, isEven)) // {0, 2, 4, 6, 8}
}
```
