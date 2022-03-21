---
title: "Drop"
date: 2022-03-21T13:48:21-04:00
---

`Drop` returns the rest of the elements in a slice. Pass an index to return the values of the slice from that index
onward.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}

	fmt.Println(u.Drop(nums, 3)) // {1, 9, 2, 3, 7, 4, 6, 5}
}
```
