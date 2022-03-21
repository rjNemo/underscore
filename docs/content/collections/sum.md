---
title: "Sum"
date: 2022-03-21T13:50:29-04:00
---

`Sum` adds elements of the slice.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(u.Sum(nums)) // 45
}
```
