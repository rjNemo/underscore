---
title: "First"
date: 2025-01-16T00:00:00-00:00
---

`First` returns the first element of the slice. Returns an error if the slice is empty.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 2, 3, 4, 5}
 first, err := u.First(nums)
 if err != nil {
  panic(err)
 }
 fmt.Println(first) // 1

 // Handle empty slice
 empty := []int{}
 _, err = u.First(empty)
 if err != nil {
  fmt.Println("Error:", err) // Error: underscore: empty slice
 }
}
```
