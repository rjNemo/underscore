---
title: "Values"
date: 2025-09-01T00:00:00-00:00
---

`maps.Values` returns the values of a map in unspecified order.

```go
package main

import (
 "fmt"
 m "github.com/rjNemo/underscore/maps"
)

func main() {
 fmt.Println(m.Values(map[int]string{1:"a",2:"b"})) // e.g., ["b" "a"]
}
```
