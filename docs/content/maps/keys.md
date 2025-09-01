---
title: "Keys"
date: 2025-09-01T00:00:00-00:00
---

`maps.Keys` returns the keys of a map in unspecified order.

```go
package main

import (
 "fmt"
 m "github.com/rjNemo/underscore/maps"
)

func main() {
 fmt.Println(m.Keys(map[int]string{1:"a",2:"b"})) // e.g., [2 1]
}
```
