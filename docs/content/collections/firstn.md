---
title: "FirstN"
date: 2025-01-16T00:00:00-00:00
---

`FirstN` returns the first n elements of the slice. If n is greater than the slice length, returns the entire slice. If n is less than or equal to 0, returns an empty slice.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
 fmt.Println(u.FirstN(nums, 3))  // [1, 2, 3]
 fmt.Println(u.FirstN(nums, 0))  // []
 fmt.Println(u.FirstN(nums, 10)) // [1, 2, 3, 4, 5, 6, 7, 8, 9]
 fmt.Println(u.FirstN(nums, -5)) // []
}
```
