# _Underscore

![License](https://img.shields.io/github/license/rjNemo/underscore?style=for-the-badge)
[![Go version](https://img.shields.io/github/go-mod/go-version/rjNemo/underscore?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/rjNemo/underscore)
![Go report](https://goreportcard.com/badge/github.com/rjNemo/underscore?style=for-the-badge)
![test coverage](https://img.shields.io/codecov/c/github/rjNemo/underscore?style=for-the-badge&logo=codecov)

![underscore](https://socialify.git.ci/rjNemo/underscore/image?description=1&font=KoHo&language=1&logo=https%3A%2F%2Fraw.githubusercontent.com%2FrjNemo%2Funderscore%2Fmain%2Fdocs%2Fstatic%2Flogo.png&owner=1&pattern=Floating%20Cogs&stargazers=1&theme=Dark)

`underscore` is a `Go` library providing useful functional programming helpers without extending any built-in objects.

It is mostly a port from the `underscore.js` library based on generics brought by `Go 1.18`.

## Usage

📚 Follow this link for the [documentation](https://underscore.onrender.com/).

Install the library using

```shell
go get github.com/rjNemo/underscore@0.4.0
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

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

You need at least `go1.18` for development. The project is shipped with a [Dockerfile](Dockerfile) based on `go1.18`.

If you prefer local development, navigate to the [official
download page](https://go.dev/dl/) and install version `1.18` or beyond.

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

## Functions

`underscore` provides many of functions that support your favorite functional helpers

### Collections

- `All`
- `Any`
- `Contains` (only numerics values at the moment)
- `Each`
- `Filter`
- `Flatmap`
- `Find`
- `Map`
- `Max`
- `Min`
- `Partition`
- `Reduce`

### Pipe

Calling `NewPipe` will cause all future method calls to return wrapped values. When you've finished the computation,
call `Value` to retrieve the final value.

Methods not returning a slice such as `Reduce`, `All`, `Any`, will break the `Chain` and return `Value` instantly.

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
