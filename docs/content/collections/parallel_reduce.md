---
title: "ParallelReduce"
date: 2025-01-16T00:00:00-00:00
---

`ParallelReduce` applies a reduction function in parallel using a worker pool. The operation must be associative and commutative for correct results. If workers <= 0, defaults to GOMAXPROCS. On error, the first error is returned and processing is canceled.

**Note:** This is an experimental function. Order of operations is not guaranteed, so use only with associative and commutative operations (like addition, multiplication, min, max).

```go
package main

import (
	"context"
	"fmt"
	"time"
	u "github.com/rjNemo/underscore"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ctx := context.Background()

	// Parallel sum (safe - addition is associative and commutative)
	result, err := u.ParallelReduce(ctx, nums, 4, func(ctx context.Context, n int, acc int) (int, error) {
		// Simulate expensive computation
		time.Sleep(10 * time.Millisecond)
		return n + acc, nil
	}, 0)

	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Result will vary due to parallel execution

	// With context cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err = u.ParallelReduce(ctx, nums, 4, func(ctx context.Context, n int, acc int) (int, error) {
		time.Sleep(100 * time.Millisecond)
		return n + acc, nil
	}, 0)

	if err != nil {
		fmt.Println("Operation was cancelled:", err)
	}
}
```

**Warning:** Do not use ParallelReduce for non-associative operations like subtraction or division, as the results will be unpredictable due to parallel execution order.
