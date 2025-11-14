package underscore_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestParallelMap_OrderAndResult(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	out, err := u.ParallelMap(context.Background(), values, 2, func(_ context.Context, n int) (int, error) {
		return n * n, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 4, 9, 16, 25}, out)
}

func TestParallelMap_Error(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	wantErr := errors.New("boom")
	out, err := u.ParallelMap(context.Background(), values, 4, func(_ context.Context, n int) (int, error) {
		if n == 3 {
			return 0, wantErr
		}
		return n, nil
	})
	assert.Error(t, err)
	assert.Nil(t, out)
}

func TestParallelMap_DefaultWorkers(t *testing.T) {
	values := []int{1, 2, 3}
	out, err := u.ParallelMap(context.Background(), values, 0, func(_ context.Context, n int) (int, error) {
		return n + 1, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 3, 4}, out)
}

func BenchmarkParallelMap(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	ctx := context.Background()

	for _, workers := range []int{1, 2, 4, 8} {
		b.Run(fmt.Sprintf("workers=%d", workers), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.ParallelMap(ctx, data, workers, func(_ context.Context, n int) (int, error) {
					return n * 2, nil
				})
			}
		})
	}
}

func BenchmarkMapVsParallelMap(b *testing.B) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}
	ctx := context.Background()

	b.Run("Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			u.Map(data, func(n int) int { return n * 2 })
		}
	})

	b.Run("ParallelMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			u.ParallelMap(ctx, data, 0, func(_ context.Context, n int) (int, error) {
				return n * 2, nil
			})
		}
	})
}
