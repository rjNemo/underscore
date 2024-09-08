---
title: "Unique"
date: 2022-04-12T17:18:04-04:00
---

`Unique` returns a duplicate-free version of the slice. Only the first occurrence of each value is kept.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 4, 2, 5, 3, 1, 5, 2, 8, 9}

 fmt.Println(u.Unique(nums)) // 1, 4, 2, 5, 3, 8, 9
}
```
