---
title: "NewPipe"
date: 2021-12-31T13:11:41-04:00
---

Calling `NewPipe` will cause all future method calls to return wrapped objects.
When you've finished the computation, call `Value` to retrieve the final value.

Methods not returning a collection such as `Reduce`, `All`, `Any`, will break the
chain and return `Value` instantly.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 sum := u.NewPipe([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).
  // filter even numbers from the slice
  Filter(func(n int) bool { return n%2 == 0 }).
  // square every number in the slice
  Map(func(n int) int { return n * n }).
  // reduce to the sum
  Reduce(func(n, acc int) int { return n + acc }, 0)

 fmt.Println(sum) // 120
}
```

