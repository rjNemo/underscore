---
title: "Any"
date: 2022-03-21T13:26:01-04:00
---

`Any` returns true if any of the values in the slice pass the predicate truth test. Short-circuits and stops traversing
the slice if a true element is found.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 2, 4, 6, 8}
 isEven := func(n int) bool { return n%2 == 0 }
 fmt.Println(u.Any(nums, isEven)) // true
}
```
