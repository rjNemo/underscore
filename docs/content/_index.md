---
title: underscore
---

![License](https://img.shields.io/github/license/rjNemo/underscore?style=for-the-badge)
[![Go version](https://img.shields.io/github/go-mod/go-version/rjNemo/underscore?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/rjNemo/underscore)
![Go report](https://goreportcard.com/badge/github.com/rjNemo/underscore?style=for-the-badge)
![test coverage](https://img.shields.io/codecov/c/github/rjNemo/underscore?style=for-the-badge&logo=codecov)

![underscore](https://socialify.git.ci/rjNemo/underscore/image?description=1&font=KoHo&language=1&logo=https%3A%2F%2Fraw.githubusercontent.com%2FrjNemo%2Funderscore%2Fmain%2Fdocs%2Fstatic%2Flogo.png&owner=1&pattern=Floating%20Cogs&stargazers=1&theme=Dark)


`underscore` is a `Go` library providing useful functional programming helpers without extending any built-in objects.

It is mostly a port from the `underscore.js` library based on generics brought by `go1.18`.

## Quick Start

Install the library using

```shell
go get github.com/rjNemo/underscore
```

Please check out the [examples](https://github.com/rjNemo/underscore/tree/main/examples) to see how to use the library.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// filter even numbers from the slice
	evens := u.Filter(numbers, func(n int) bool { return n%2 == 0 })
	// square every number in the slice
	squares := u.Map(evens, func(n int) int { return n * n })
	// reduce to the sum
	res := u.Reduce(squares, func(n, acc int) int { return n + acc }, 0)

	fmt.Println(res) // 120
}
```
