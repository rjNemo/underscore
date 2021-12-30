# Underscore

![License](https://img.shields.io/github/license/rjNemo/underscore?style=for-the-badge)
![Go version](https://img.shields.io/github/go-mod/go-version/rjNemo/underscore?style=for-the-badge)

![underscore](https://socialify.git.ci/rjNemo/underscore/image?description=1&descriptionEditable=A%20library%20providing%20useful%20functional%20programming%20helpers%20for%20Go%201.18&font=Raleway&language=1&logo=https%3A%2F%2Fgithub.com%2FrjNemo%2Funderscore%2Fblob%2Fmain%2Fdocs%2Funderscore.png%3Fraw%3Dtrue&name=1&pattern=Floating%20Cogs&theme=Light)

`underscore` is a `Go` library providing useful functional programming helpers without extending any built-in objects.

It is mostly a port from the `underscore.js` library based on generics brought by `go1.18`.

## Usage

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
	isEven := func(n int) bool { return n%2 == 0 }
	evens := u.Filter(numbers, isEven)
	// square every number in the slice
	toSquare := func(n int) int { return n * n }
	squares := u.Map(evens, toSquare)
	// reduce to the sum 
	sum := func(n, acc int) int { return n + acc }
	res := u.Reduce(squares, sum, 0)

	fmt.Println(res) // 110
}
```

## Getting Started

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

Building the docker image will run the tests automatically. Otherwise, you can simply run:

```shell
go test ./...
```

## Functions

`underscore` provides 100s of functions that support your favorite functional helpers

### Collections

- `map`
- `filter`
- `reduce`
- `each`
- `some`
- `every`
- `find`
- `contains` (only numerics values at the moment)
- `max`
- `min`
- `partition`

## Built With

- [Go](https://go.dev/) - Build fast, reliable, and efficient software at scale

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull
requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see
the [tags on this repository](https://github.com/rjNemo/underscore/tags).

## Authors

- **Ruidy** - _Initial work_ - [Ruidy](https://github.com/rjNemo)

See also the list of [contributors](https://github.com/rjNemo/underscore/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

This project is largely inspired by [Underscore.js](https://underscorejs.org/#) library. Check out the original project
if you don't already know it.