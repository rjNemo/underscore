package underscore

import (
	"context"
	"runtime"
	"sync"
)

// ParallelFilter filters values using a context-aware predicate concurrently and preserves input order.
// If workers <= 0, it defaults to GOMAXPROCS. On error, cancels work and returns nil with the error.
func ParallelFilter[T any](ctx context.Context, values []T, workers int, fn func(context.Context, T) (bool, error)) ([]T, error) {
	if workers <= 0 {
		workers = runtime.GOMAXPROCS(0)
	}
	type task struct {
		idx int
		val T
	}

	keeps := make([]bool, len(values))
	tasks := make(chan task)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	var once sync.Once
	var firstErr error

	worker := func() {
		defer wg.Done()
		for t := range tasks {
			select {
			case <-ctx.Done():
				return
			default:
			}
			keep, err := fn(ctx, t.val)
			if err != nil {
				once.Do(func() {
					firstErr = err
					cancel()
				})
				continue
			}
			keeps[t.idx] = keep
		}
	}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker()
	}

OUTER:
	for i, v := range values {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			tasks <- task{idx: i, val: v}
		}
	}
	close(tasks)
	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}

	// Build result preserving order
	// Pre-count capacity to avoid re-allocations
	count := 0
	for _, k := range keeps {
		if k {
			count++
		}
	}
	res := make([]T, 0, count)
	for i, k := range keeps {
		if k {
			res = append(res, values[i])
		}
	}
	return res, nil
}
