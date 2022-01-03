---
title: "Collections"
date: 2021-12-30T13:24:39-04:00
---
 
## All

`All` returns true if all the values in the slice pass the predicate truth test.\
Short-circuits and stops traversing the slice if a false element is found.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 3, 5, 7, 9}
	isOdd := func(n int) bool { return n%2 != 0 }
	fmt.Println(u.All(nums, isOdd)) // true
}
```

## Contains

`Contains` returns true if the value is present in the slice.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 3, 5, 7, 9}

	fmt.Println(u.Contains(nums, 5)) // true 
}
```

## Each

`Each` iterates over a slice of elements, yielding each in turn to an action function.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	names := []string{"Alice", "Bob", "Charles"}
	res := make([]string, 0)

	u.Each(names, func(n string) {
		res = append(res, fmt.Sprintf("Hi %s", n))
	})
	fmt.Println(res) // {"Hi Alice", "Hi Bob", "Hi Charles"} 
}
```

## Filter

`Filter` looks through each value in the slice, returning a slice of all the values that pass a truth test (predicate).

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(n int) bool { return n%2 == 0 }
	fmt.Println(u.Filter(nums, isEven)) // {0, 2, 4, 6, 8}
}
```

## Find

Find looks through each value in the slice, returning the first one that passes a truth test (predicate), or the default
value for the type and an error if no value passes the test. The function returns as soon as it finds an acceptable
element, and doesn't traverse the entire slice.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{2, 4, 5, 6, 8, 0}
	isOdd := func(n int) bool { return n%2 != 0 }

	n, err := u.Find(nums, isOdd)
	fmt.Println(n)   // 5
	fmt.Println(err) // nil
}
```

## Map

`Map` produces a new slice of values by mapping each value in the slice through a transform function.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 2, 3}
	toSquare := func(n int) int {
		return n * n
	}
	fmt.Println(u.Map(nums, toSquare)) // {1, 4, 9}
}
```

## Max

`Max` returns the maximum value in the slice. This function can currently only compare numbers reliably. This function
uses operator `<`.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	fmt.Println(u.Max(nums)) // 9
}
```

## Min

`Min` returns the minimum value in the slice. This function can currently only compare numbers reliably. This function
uses operator `<`.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	fmt.Println(u.Min(nums)) // 1
}
```

## Partition

`Partition` splits the slice into two slices: one whose elements all satisfy predicate and one whose elements all do not
satisfy predicate.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(n int) bool { return n%2 == 0 }

	evens, odds := u.Partition(nums, isEven)
	fmt.Println(evens) // {0, 2, 4, 6, 8}
	fmt.Println(odds)  // {1, 3, 5, 7, 9}
}
```

## Reduce

`Reduce` combine a list of values into a single value. `acc` is the initial state, and each successive step of it should
be returned by the reduction function.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := func(n, acc int) int { return n + acc }
	fmt.Println(u.Reduce(nums, sum, 0)) // 45
}
```

## Some

`Some` returns true if any of the values in the slice pass the predicate truth test. Short-circuits and stops traversing
the slice if a true element is found.

```go
package main

import (
	"fmt"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 2, 4, 6, 8}
	isEven := func(n int) bool { return n%2 == 0 }
	fmt.Println(u.Some(nums, isEven)) // true
}
```