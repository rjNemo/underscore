---
title: "ContainsBy"
date: 2025-09-01T00:00:00-00:00
---

`ContainsBy` returns true if any element satisfies the predicate.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 3, 5, 8}
 fmt.Println(u.ContainsBy(nums, func(n int) bool { return n%2 == 0 })) // true
}
```
