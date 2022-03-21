---
title: "Partition"
date: 2022-03-21T13:33:23-04:00
---

`Partition` splits the slice into two slices: one whose elements all satisfy predicate and one whose elements all do not
satisfy predicate.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(n int) bool { return n%2 == 0 }

	evens, odds := u.Partition(nums, isEven)
	fmt.Println(evens) // {0, 2, 4, 6, 8}
	fmt.Println(odds)  // {1, 3, 5, 7, 9}
}
```
