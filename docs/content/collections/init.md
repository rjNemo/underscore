---
title: "Init"
date: 2025-01-16T00:00:00-00:00
---

`Init` returns all elements except the last one, and the last element separately. Returns an empty slice and zero value if the input slice is empty. Useful for destructuring lists from the right.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	init, last := u.Init(nums)
	fmt.Println(init) // [1, 2, 3, 4]
	fmt.Println(last) // 5

	// Single element
	single, val := u.Init([]int{42})
	fmt.Println(single) // []
	fmt.Println(val)    // 42

	// Empty slice
	empty, zero := u.Init([]int{})
	fmt.Println(empty) // []
	fmt.Println(zero)  // 0
}
```
