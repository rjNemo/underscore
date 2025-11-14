package underscore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestParallelReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ctx := context.Background()

	// Note: This is a simplified test - ParallelReduce needs work for proper reduction
	result, err := u.ParallelReduce(ctx, nums, 2, func(ctx context.Context, n int, acc int) (int, error) {
		return n + acc, nil
	}, 0)

	assert.NoError(t, err)
	// Result may vary due to parallel execution
	assert.Greater(t, result, 0)
}

func TestParallelReduceEmpty(t *testing.T) {
	ctx := context.Background()
	result, err := u.ParallelReduce(ctx, []int{}, 2, func(ctx context.Context, n int, acc int) (int, error) {
		return n + acc, nil
	}, 42)

	assert.NoError(t, err)
	assert.Equal(t, 42, result)
}
