---
title: "Chunk"
date: 2025-09-01T00:00:00-00:00
---

`Chunk` splits a slice into groups of size `n`. The last chunk may be smaller.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 fmt.Println(u.Chunk([]int{1,2,3,4,5}, 2)) // [[1 2] [3 4] [5]]
}
```
