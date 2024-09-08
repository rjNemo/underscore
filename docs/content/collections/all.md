---
title: "All"
date: 2021-12-30T13:24:39-04:00
---

`All` returns true if all the values in the slice pass the predicate truth test.\
Short-circuits and stops traversing the slice if a false element is found.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 3, 5, 7, 9}
 isOdd := func(n int) bool { return n%2 != 0 }
 fmt.Println(u.All(nums, isOdd)) // true
}
```
