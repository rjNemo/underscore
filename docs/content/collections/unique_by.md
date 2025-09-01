---
title: "UniqueBy"
date: 2025-09-01T00:00:00-00:00
---

`UniqueBy` returns a duplicate-free version of the slice using a key selector.
Order is preserved; the first occurrence of each key is kept.

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

type User struct{ ID int; Email string }

func main() {
 users := []User{{1, "a@x"}, {2, "b@x"}, {3, "a@x"}}
 fmt.Println(u.UniqueBy(users, func(u User) string { return u.Email }))
 // [{1 a@x} {2 b@x}]
}
```
