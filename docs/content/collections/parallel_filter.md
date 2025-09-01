---
title: "ParallelFilter"
date: 2025-09-01T00:00:00-00:00
---

`ParallelFilter` filters a slice concurrently with a worker pool, preserves order,
and supports context cancellation.

```go
package main

import (
 "context"
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 out, err := u.ParallelFilter(context.Background(), []int{1,2,3,4,5}, 3,
  func(ctx context.Context, n int) (bool, error) { return n%2==0, nil },
 )
 fmt.Println(out, err) // [2 4] <nil>
}
```
