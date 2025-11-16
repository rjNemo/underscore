---
title: "Sliding"
date: 2025-01-16T00:00:00-00:00
---

`Sliding` creates a sliding window view of the slice with the specified window size. Returns an empty slice if size is less than or equal to 0 or greater than the slice length. Useful for moving averages, n-grams, and pattern matching.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 2, 3, 4, 5}
 fmt.Println(u.Sliding(nums, 3)) // [[1, 2, 3], [2, 3, 4], [3, 4, 5]]

 // Size 2
 fmt.Println(u.Sliding(nums, 2)) // [[1, 2], [2, 3], [3, 4], [4, 5]]

 // N-grams for text
 words := []string{"the", "quick", "brown", "fox"}
 bigrams := u.Sliding(words, 2)
 fmt.Println(bigrams) // [["the", "quick"], ["quick", "brown"], ["brown", "fox"]]

 // Moving average example
 data := []int{10, 20, 30, 40, 50}
 windows := u.Sliding(data, 3)
 for _, window := range windows {
  sum := 0
  for _, v := range window {
   sum += v
  }
  avg := sum / len(window)
  fmt.Printf("Window: %v, Average: %d\n", window, avg)
 }
 // Window: [10 20 30], Average: 20
 // Window: [20 30 40], Average: 30
 // Window: [30 40 50], Average: 40
}
```
