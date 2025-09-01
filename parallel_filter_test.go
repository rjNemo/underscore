package underscore_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestParallelFilter_OrderAndResult(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	out, err := u.ParallelFilter(context.Background(), values, 3, func(_ context.Context, n int) (bool, error) {
		return n%2 == 0, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4}, out)
}

func TestParallelFilter_Error(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	boom := errors.New("boom")
	out, err := u.ParallelFilter(context.Background(), values, 2, func(_ context.Context, n int) (bool, error) {
		if n == 4 {
			return false, boom
		}
		return true, nil
	})
	assert.Error(t, err)
	assert.Nil(t, out)
}
