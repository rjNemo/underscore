---
title: "Reduce"
date: 2022-03-21T13:33:52-04:00
---

`Reduce` combine a list of values into a single value. `acc` is the initial state, and each successive step of it should
be returned by the reduction function.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
 sum := func(n, acc int) int { return n + acc }
 fmt.Println(u.Reduce(nums, sum, 0)) // 45
}
```
