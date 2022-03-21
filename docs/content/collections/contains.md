---
title: "Contains"
date: 2022-03-21T13:30:29-04:00
---

`Contains` returns true if the value is present in the slice.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 3, 5, 7, 9}

	fmt.Println(u.Contains(nums, 5)) // true 
}
```
