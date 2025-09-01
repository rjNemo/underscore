---
title: "UniqueInPlace"
date: 2025-09-01T00:00:00-00:00
---

`UniqueInPlace` removes duplicates from a slice in place while preserving order.
Returns the shortened slice.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 xs := []int{1,4,2,5,3,1,5,2}
 fmt.Println(u.UniqueInPlace(xs)) // [1 4 2 5 3]
}
```
