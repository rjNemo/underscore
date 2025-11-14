package underscore_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

// Large data stress tests

func TestFilterLargeData(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	large := make([]int, 1_000_000)
	for i := range large {
		large[i] = i
	}

	result := u.Filter(large, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, 500_000, len(result))
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 999_998, result[len(result)-1])
}

func TestMapLargeData(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	large := make([]int, 1_000_000)
	for i := range large {
		large[i] = i
	}

	result := u.Map(large, func(n int) int { return n * 2 })
	assert.Equal(t, 1_000_000, len(result))
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 1_999_998, result[len(result)-1])
}

func TestPartitionLargeData(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	large := make([]int, 1_000_000)
	for i := range large {
		large[i] = i
	}

	keep, reject := u.Partition(large, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, 500_000, len(keep))
	assert.Equal(t, 500_000, len(reject))
}

func TestUniqueLargeData(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	large := make([]int, 1_000_000)
	for i := range large {
		large[i] = i % 1000 // Many duplicates
	}

	result := u.Unique(large)
	assert.Equal(t, 1000, len(result))
}

// Concurrency stress tests

func TestParallelMapHighConcurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}

	ctx := context.Background()

	// Test with many workers
	result, err := u.ParallelMap(ctx, data, 100, func(ctx context.Context, n int) (int, error) {
		time.Sleep(time.Microsecond) // Simulate work
		return n * 2, nil
	})

	assert.NoError(t, err)
	assert.Equal(t, len(data), len(result))
	for i, v := range result {
		assert.Equal(t, data[i]*2, v)
	}
}

func TestParallelMapCancellation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	_, err := u.ParallelMap(ctx, data, 4, func(ctx context.Context, n int) (int, error) {
		// Check context and return error if canceled
		if ctx.Err() != nil {
			return 0, ctx.Err()
		}
		time.Sleep(1 * time.Millisecond) // Slow work
		return n, nil
	})

	// Should either complete or return a context error
	if err != nil {
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	}
}

func TestParallelFilterHighConcurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}

	ctx := context.Background()

	result, err := u.ParallelFilter(ctx, data, 50, func(ctx context.Context, n int) (bool, error) {
		time.Sleep(time.Microsecond)
		return n%2 == 0, nil
	})

	assert.NoError(t, err)
	assert.Equal(t, 5000, len(result))
	for _, v := range result {
		assert.Equal(t, 0, v%2)
	}
}

// Race condition tests

func TestParallelMapNoRaces(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	// Run with: go test -race -run TestParallelMapNoRaces
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	ctx := context.Background()

	for i := 0; i < 100; i++ {
		_, err := u.ParallelMap(ctx, data, 8, func(ctx context.Context, n int) (int, error) {
			return n * 2, nil
		})
		assert.NoError(t, err)
	}
}

func TestParallelFilterNoRaces(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	// Run with: go test -race -run TestParallelFilterNoRaces
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	ctx := context.Background()

	for i := 0; i < 100; i++ {
		_, err := u.ParallelFilter(ctx, data, 8, func(ctx context.Context, n int) (bool, error) {
			return n%2 == 0, nil
		})
		assert.NoError(t, err)
	}
}

func TestConcurrentFilterCalls(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	// Test that concurrent calls to Filter don't interfere with each other
	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}

	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			result := u.Filter(data, func(n int) bool { return n%2 == 0 })
			if len(result) != 5000 {
				t.Errorf("Expected 5000 elements, got %d", len(result))
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestConcurrentMapCalls(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	data := make([]int, 10000)
	for i := range data {
		data[i] = i
	}

	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			result := u.Map(data, func(n int) int { return n * 2 })
			if len(result) != 10000 {
				t.Errorf("Expected 10000 elements, got %d", len(result))
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
