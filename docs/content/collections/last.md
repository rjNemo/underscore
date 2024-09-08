---
title: "Last"
date: 2022-03-21T13:46:24-04:00
---

`Last` returns the last element of the slice.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}

 fmt.Println(u.Last(nums)) // 5
}
```
