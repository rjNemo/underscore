---
title: "Unzip"
date: 2025-01-16T00:00:00-00:00
---

`Unzip` splits a slice of tuples into two separate slices. The inverse operation of Zip. Useful for separating paired data.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 // Basic usage
 pairs := []u.Tuple[int, string]{
  {Left: 1, Right: "a"},
  {Left: 2, Right: "b"},
  {Left: 3, Right: "c"},
 }

 nums, letters := u.Unzip(pairs)
 fmt.Println(nums)    // [1, 2, 3]
 fmt.Println(letters) // ["a", "b", "c"]

 // Use case: separating keys and values
 keyValuePairs := []u.Tuple[string, int]{
  {Left: "apple", Right: 5},
  {Left: "banana", Right: 3},
  {Left: "cherry", Right: 8},
 }

 items, counts := u.Unzip(keyValuePairs)
 fmt.Println("Items:", items)   // Items: [apple banana cherry]
 fmt.Println("Counts:", counts) // Counts: [5 3 8]

 // Empty slice
 emptyNums, emptyStrs := u.Unzip([]u.Tuple[int, string]{})
 fmt.Println(emptyNums, emptyStrs) // [] []
}
```
