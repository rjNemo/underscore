# \_Underscore

![License](https://img.shields.io/github/license/rjNemo/underscore?style=for-the-badge)
[![Go version](https://img.shields.io/github/go-mod/go-version/rjNemo/underscore?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/rjNemo/underscore)
![Go report](https://goreportcard.com/badge/github.com/rjNemo/underscore?style=for-the-badge)
![test coverage](https://img.shields.io/codecov/c/github/rjNemo/underscore?style=for-the-badge&logo=codecov)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/9726/badge?style=for-the-badge)](https://www.bestpractices.dev/projects/9726)

![underscore](https://socialify.git.ci/rjNemo/underscore/image?description=1&font=KoHo&language=1&logo=https%3A%2F%2Fraw.githubusercontent.com%2FrjNemo%2Funderscore%2Fmain%2Fdocs%2Fstatic%2Flogo.png&owner=1&pattern=Floating%20Cogs&stargazers=1&theme=Dark)

`underscore` is a `Go` library providing useful functional programming helpers without
extending any built-in objects.

It is mostly a port from the `underscore.js` library based on generics brought by
`Go 1.18`.

## Usage

📚 Follow this link for the [documentation](https://underscore.onrender.com/).

Install the library using

```sh
go get github.com/rjNemo/underscore@latest
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

These instructions will get you a copy of the project up and running on your local
machine for development and testing purposes.

### Prerequisites

You need at least `go1.24` for development. The project is shipped with a [Dockerfile](Dockerfile)
based on `go1.24`.

If you prefer local development, navigate to the [official
download page](https://go.dev/dl/) and install version `1.24` or beyond.

### Installing

First clone the repository

```sh
git clone https://github.com/rjNemo/underscore.git
```

Install dependencies

```sh
go mod download
```

And that's it.

## Tests

To run the unit tests, you can simply run:

```sh
make test
```

## Functions

`underscore` provides many of functions that support your favorite functional helpers

### Collections

- `All`
- `Any`
- `Chunk`
- `Contains`
- `ContainsBy`
- `Count`
- `Difference`
- `Drop`
- `Each`
- `Filter`
- `Find`
- `Flatmap`
- `GroupBy`
- `Intersection`
- `Join` / `JoinProject`
- `Last`
- `Map`
- `Max`
- `Min`
- `OrderBy`
- `Partition`
- `Range`
- `Reduce`
- `RemoveAt`
- `Sum` / `SumMap`
- `Unique`
- `UniqueBy`
- `UniqueInPlace`
- `Zip`

### Pipe

Calling `NewPipe` will cause all future method calls to return wrapped values. When
you've finished the computation, call `Value` to retrieve the final value.

Methods not returning a slice such as `Reduce`, `All`, `Any`, will break the `Chain`
and return `Value` instantly.

### Concurrency

- `ParallelMap(ctx, values, workers, fn)`: apply a function concurrently while preserving order and supporting context cancellation.
- `ParallelFilter(ctx, values, workers, fn)`: filter concurrently with order preserved and context support.

```go
package main

import (
 "context"
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 out, err := u.ParallelMap(context.Background(), []int{1, 2, 3, 4}, 4,
  func(ctx context.Context, n int) (int, error) { return n * n, nil },
 )
 fmt.Println(out, err) // [1 4 9 16] <nil>
}
```

```go
// ParallelFilter example
package main

import (
 "context"
 "fmt"
 u "github.com/rjNemo/underscore"
)

func main() {
 out, err := u.ParallelFilter(context.Background(), []int{1,2,3,4,5}, 3,
  func(ctx context.Context, n int) (bool, error) { return n%2==0, nil },
 )
 fmt.Println(out, err) // [2 4] <nil>
}
```

### Utilities

- `Ternary`: conditional expression helper
- `ToPointer`: convert values to pointers
- `SortSliceASC` / `SortSliceDESC`: sort slices in ascending or descending order
- `Result`, `Ok`, `Err`, `ToResult`: Result type for error handling
- `Tuple`: generic tuple type for paired values

### Subpackages

- `maps.Keys(m)` / `maps.Values(m)`: extract keys or values from maps
- `maps.Map(m, fn)`: transform map entries

## Built With

- [Go](https://go.dev/) - Build fast, reliable, and efficient software at scale

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct,
and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see
the [tags on this repository](https://github.com/rjNemo/underscore/tags).

## Authors

- **Ruidy** - _Initial work_ - [Ruidy](https://github.com/rjNemo)

See also the list of [contributors](https://github.com/rjNemo/underscore/contributors)
who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md)
file for details

## Acknowledgments

This project is largely inspired by [Underscore.js](https://underscorejs.org/#)
library. Check out the original project if you don't already know it.
