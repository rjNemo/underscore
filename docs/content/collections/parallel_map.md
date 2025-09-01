---
title: "ParallelMap"
date: 2025-09-01T00:00:00-00:00
---

`ParallelMap` applies a function to each element concurrently using a worker pool,
preserves order, and supports context cancellation.

```go
package main

import (
 "context"
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 out, err := u.ParallelMap(context.Background(),
    []int{1,2,3,4}, 4, func(ctx context.Context, n int) (int, error) {
  return n*n, nil
 })
 fmt.Println(out, err) // [1 4 9 16] <nil>
}
```
