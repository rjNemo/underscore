---
title: "Map"
date: 2022-03-21T13:32:10-04:00
---

`Map` produces a new slice of values by mapping each value in the slice through a transform function.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 2, 3}
 toSquare := func(n int) int {
  return n * n
 }
 fmt.Println(u.Map(nums, toSquare)) // {1, 4, 9}
}
```
