---
title: "DropWhile"
date: 2025-01-16T00:00:00-00:00
---

`DropWhile` drops elements from the beginning of the slice while the predicate returns true. It returns the remaining elements starting from the first element where the predicate returns false.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
 lessThan5 := func(n int) bool { return n < 5 }
 fmt.Println(u.DropWhile(nums, lessThan5)) // [5, 6, 7, 8, 9]

 words := []string{"apple", "banana", "cherry", "date"}
 shortWords := func(s string) bool { return len(s) < 6 }
 fmt.Println(u.DropWhile(words, shortWords)) // ["banana", "cherry", "date"]
}
```
