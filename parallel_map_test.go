package underscore_test

import (
	"context"
	"errors"
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
