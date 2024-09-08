---
title: "Flatmap"
date: 2022-08-10T16:49:56+02:00
draft: false
---

Flatmap flattens the input slice element into the new slice. FlatMap maps every element with the help of a mapper function, then flattens the input slice element into the new slice.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
  nums := []int{1, 2, 3, 4}
  mapper := func(n int) []int { return []int{(n - 1) * n, (n) * n} }
  res := u.Flatmap(nums, mapper)
  fmt.Println(res) // {0, 1, 2, 4, 6, 9, 12, 16}
}
```
