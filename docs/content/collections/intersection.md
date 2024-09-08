---
title: "Intersection"
date: 2022-03-27T00:24:11-04:00
---

Intersection computes the list of values that are the intersection of all the slices.
Each value in the result is present in each of the slices.

```go
package main


import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main(){
    a := []int{1, 3, 5, 7, 9}
    b := []int{2, 3, 5, 8, 0}

    fmt.Println(u.Intersection(a, b)) // {3, 5}
}
```

