---
title: "Find"
date: 2022-03-21T13:31:40-04:00
---

Find looks through each value in the slice, returning the first one that passes a truth test (predicate), or the default
value for the type and an error if no value passes the test. The function returns as soon as it finds an acceptable
element, and doesn't traverse the entire slice.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{2, 4, 5, 6, 8, 0}
 isOdd := func(n int) bool { return n%2 != 0 }

 n, err := u.Find(nums, isOdd)
 fmt.Println(n)   // 5
 fmt.Println(err) // nil
}
```
