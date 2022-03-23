---
title: "Each"
date: 2022-03-21T13:30:59-04:00
---

`Each` iterates over a slice of elements, yielding each in turn to an action function.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	names := []string{"Alice", "Bob", "Charles"}
	res := make([]string, 0)

	u.Each(names, func(n string) {
		res = append(res, fmt.Sprintf("Hi %s", n))
	})
	fmt.Println(res) // {"Hi Alice", "Hi Bob", "Hi Charles"} 
}
```
