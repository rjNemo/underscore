---
title: "Intersperse"
date: 2025-01-16T00:00:00-00:00
---

`Intersperse` inserts a separator between each element of the slice. Returns an empty slice if the input is empty. Returns the original element if the input has only one element.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 2, 3, 4, 5}
 fmt.Println(u.Intersperse(nums, 0)) // [1, 0, 2, 0, 3, 0, 4, 0, 5]

 // Useful for formatting
 words := []string{"apple", "banana", "cherry"}
 fmt.Println(u.Intersperse(words, ",")) // ["apple", ",", "banana", ",", "cherry"]

 // Single element - no separator added
 single := []int{42}
 fmt.Println(u.Intersperse(single, 0)) // [42]
}
```
