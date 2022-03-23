---
title: "Min"
date: 2022-03-21T13:33:03-04:00
---

`Min` returns the minimum value in the slice. This function can currently only compare numbers reliably. This function
uses operator `<`.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	fmt.Println(u.Min(nums)) // 1
}
```
