---
title: "Getting Started"
date: 2022-03-21T14:09:01-04:00
---

## Quick Start

Install the library using

```shell
go get github.com/rjNemo/underscore
```

Please check out the [examples](examples) to see how to use the library.

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

## Installation

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

You need at least `go1.18` for development. The project is shipped with a [Dockerfile](Dockerfile) based on `go1.18`. If
you prefer local development, at the moment the easiest way to do it:

```shell script
go install golang.org/dl/go1.18beta1@latest
go1.18beta1 download
```

### Installing

First clone the repository

```shell
git clone https://github.com/rjNemo/underscore.git
```

Install dependencies

```shell
go mod download
```

And that's it.

## Tests

To run the unit tests, you can simply run:

```shell
make test
```
