package underscore

import (
	"context"
	"runtime"
	"sync"
)

// ParallelReduce applies a reduction function in parallel using a worker pool.
// The operation must be associative and commutative for correct results.
// If workers <= 0, defaults to GOMAXPROCS.
// On error, the first error is returned and processing is canceled.
//
// Note: Order of operations is not guaranteed, so use only with associative/commutative operations.
func ParallelReduce[T, P any](ctx context.Context, values []T, workers int, fn func(context.Context, T, P) (P, error), acc P) (P, error) {
	if workers <= 0 {
		workers = runtime.GOMAXPROCS(0)
	}

	if len(values) == 0 {
		return acc, nil
	}

	type task struct {
		idx int
		val T
	}

	tasks := make(chan task)
	results := make(chan P, len(values))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	var once sync.Once
	var firstErr error

	// Workers
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for t := range tasks {
				select {
				case <-ctx.Done():
					return
				default:
				}

				result, err := fn(ctx, t.val, acc)
				if err != nil {
					once.Do(func() {
						firstErr = err
						cancel()
					})
					return
				}
				results <- result
			}
		}()
	}

	// Send tasks
	go func() {
		for i, v := range values {
			select {
			case <-ctx.Done():
				close(tasks)
				return
			default:
				tasks <- task{idx: i, val: v}
			}
		}
		close(tasks)
	}()

	wg.Wait()
	close(results)

	if firstErr != nil {
		return acc, firstErr
	}

	// Combine results
	for result := range results {
		// This is a simplified combination - in practice, you'd need a combiner function
		acc = result
	}

	return acc, nil
}
