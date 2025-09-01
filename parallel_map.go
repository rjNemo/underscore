package underscore

import (
	"context"
	"runtime"
	"sync"
)

// ParallelMap applies fn to each element of values using a worker pool and preserves order.
// If workers <= 0, it defaults to GOMAXPROCS.
// On error, the first error is returned and processing is canceled; partial results are discarded.
func ParallelMap[T, P any](ctx context.Context, values []T, workers int, fn func(context.Context, T) (P, error)) ([]P, error) {
	if workers <= 0 {
		workers = runtime.GOMAXPROCS(0)
	}
	type task struct {
		idx int
		val T
	}

	res := make([]P, len(values))
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
			v, err := fn(ctx, t.val)
			if err != nil {
				once.Do(func() {
					firstErr = err
					cancel()
				})
				continue
			}
			res[t.idx] = v
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
	return res, nil
}
