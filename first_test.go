package underscore_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestFirst(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result, err := u.First(nums)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}

func TestFirstEmpty(t *testing.T) {
	_, err := u.First([]int{})
	assert.Error(t, err)
	assert.True(t, errors.Is(err, u.ErrEmptySlice))
}

func TestFirstSingleElement(t *testing.T) {
	result, err := u.First([]int{42})
	assert.NoError(t, err)
	assert.Equal(t, 42, result)
}

func TestFirstStrings(t *testing.T) {
	words := []string{"hello", "world"}
	result, err := u.First(words)
	assert.NoError(t, err)
	assert.Equal(t, "hello", result)
}

func TestFirstN(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := u.FirstN(nums, 3)
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestFirstNEmpty(t *testing.T) {
	result := u.FirstN([]int{}, 3)
	assert.Equal(t, []int{}, result)
}

func TestFirstNZero(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.FirstN(nums, 0)
	assert.Equal(t, []int{}, result)
}

func TestFirstNNegative(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.FirstN(nums, -5)
	assert.Equal(t, []int{}, result)
}

func TestFirstNGreaterThanLength(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.FirstN(nums, 10)
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestFirstNSingleElement(t *testing.T) {
	result := u.FirstN([]int{42}, 1)
	assert.Equal(t, []int{42}, result)
}

func TestFirstNAll(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.FirstN(nums, 3)
	assert.Equal(t, []int{1, 2, 3}, result)
}

func BenchmarkFirst(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = u.First(nums)
	}
}

func BenchmarkFirstN(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.FirstN(nums, 100)
	}
}
