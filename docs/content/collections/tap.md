---
title: "Tap"
date: 2025-01-16T00:00:00-00:00
---

`Tap` applies a function to each element for side effects (like debugging or logging) and returns the original slice unchanged. Useful for debugging pipelines without breaking the flow.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 // Debugging a pipeline
 nums := []int{1, 2, 3, 4, 5}

 result := u.Tap(
  u.Map(
   u.Filter(nums, func(n int) bool { return n%2 == 0 }),
   func(n int) int { return n * 2 },
  ),
  func(n int) {
   fmt.Printf("Debug: %d\n", n) // Prints each value
  },
 )

 fmt.Println(result) // [4, 8]

 // Counting elements that pass through
 count := 0
 filtered := u.Tap(
  u.Filter(nums, func(n int) bool { return n > 2 }),
  func(n int) { count++ },
 )
 fmt.Printf("Found %d elements: %v\n", count, filtered)
 // Found 3 elements: [3 4 5]

 // Logging transformations
 data := []string{"hello", "world"}
 u.Tap(data, func(s string) {
  fmt.Printf("Processing: %s\n", s)
 })
}
```
