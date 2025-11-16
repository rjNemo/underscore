---
title: "Transpose"
date: 2025-01-16T00:00:00-00:00
---

`Transpose` flips a matrix over its diagonal, swapping rows and columns. Returns an empty slice if the input is empty. Assumes all rows have the same length (uses the length of the first row).

```go
package main

import (
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 // 2x3 matrix becomes 3x2 matrix
 matrix := [][]int{
  {1, 2, 3},
  {4, 5, 6},
 }
 transposed := u.Transpose(matrix)
 fmt.Println(transposed)
 // [[1, 4], [2, 5], [3, 6]]

 // Square matrix
 square := [][]int{
  {1, 2},
  {3, 4},
 }
 fmt.Println(u.Transpose(square))
 // [[1, 3], [2, 4]]

 // Use case: converting rows to columns for processing
 data := [][]string{
  {"Name", "Age", "City"},
  {"Alice", "30", "NYC"},
  {"Bob", "25", "LA"},
 }
 byColumn := u.Transpose(data)
 fmt.Println("Names:", byColumn[0])   // [Name Alice Bob]
 fmt.Println("Ages:", byColumn[1])    // [Age 30 25]
 fmt.Println("Cities:", byColumn[2])  // [City NYC LA]
}
```
