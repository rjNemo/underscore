---
title: "Group by"
date: 2023-06-07T00:49:56+02:00
---

GroupBy splits a slice into a map[K][]V grouped by the result of the iterator function.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
  	nums := []float64{1.3, 2.1, 2.4}
	groupingFunc := func(n float64) int { return int(math.Floor(n)) }
  	res := u.GroupBy(nums, groupingFunc) // { 1: {1.3}, 2: {2.1, 2.4}}
}
```
