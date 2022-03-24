---
title: "Difference"
date: 2022-03-21T13:48:21-04:00
---

Returns a copy of the array with all instances of the values that are not present in the other arrays.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 3, 5, 6, 7, 9}
	reject := []int{9, 7, 5, 4}
	
	fmt.Println(u.Difference(nums, reject)) // {1, 3, 6}
}
```
