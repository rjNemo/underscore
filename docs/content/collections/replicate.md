---
title: "Replicate"
date: 2025-01-16T00:00:00-00:00
---

`Replicate` creates a slice containing count copies of value. Returns an empty slice if count is less than or equal to 0. Useful for initialization and testing.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 // Basic usage
 fmt.Println(u.Replicate(3, "hello"))
 // ["hello", "hello", "hello"]

 // Numbers
 fmt.Println(u.Replicate(5, 0))
 // [0, 0, 0, 0, 0]

 // Zero count
 fmt.Println(u.Replicate(0, 42))
 // []

 // Negative count
 fmt.Println(u.Replicate(-5, "x"))
 // []

 // Use case: initialize with default values
 defaultScores := u.Replicate(10, 100)
 fmt.Println(defaultScores)
 // [100, 100, 100, 100, 100, 100, 100, 100, 100, 100]

 // Use case: creating separators
 separator := u.Replicate(40, "-")
 fmt.Println(u.Reduce(separator, func(s, acc string) string { return acc + s }, ""))
 // ----------------------------------------
}
```
