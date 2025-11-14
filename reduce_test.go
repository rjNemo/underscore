package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestReduce(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reducer := func(n, acc int) int {
		return n + acc
	}
	want := 45

	assert.Equal(t, want, u.Reduce(nums, reducer, 0))
}

func TestReduceEmpty(t *testing.T) {
	result := u.Reduce([]int{}, func(n, acc int) int { return n + acc }, 10)
	assert.Equal(t, 10, result) // Should return initial accumulator
}

func TestReduceSingleElement(t *testing.T) {
	result := u.Reduce([]int{5}, func(n, acc int) int { return n + acc }, 0)
	assert.Equal(t, 5, result)
}
