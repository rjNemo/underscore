package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestSliding(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := u.Sliding(nums, 3)
	expected := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	assert.Equal(t, expected, result)
}

func TestSlidingEmpty(t *testing.T) {
	result := u.Sliding([]int{}, 3)
	assert.Equal(t, [][]int{}, result)
}

func TestSlidingSizeOne(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.Sliding(nums, 1)
	expected := [][]int{{1}, {2}, {3}}
	assert.Equal(t, expected, result)
}

func TestSlidingSizeEqualLength(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.Sliding(nums, 3)
	expected := [][]int{{1, 2, 3}}
	assert.Equal(t, expected, result)
}

func TestSlidingSizeGreaterThanLength(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.Sliding(nums, 5)
	assert.Equal(t, [][]int{}, result)
}

func TestSlidingSizeZero(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.Sliding(nums, 0)
	assert.Equal(t, [][]int{}, result)
}

func TestSlidingSizeNegative(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.Sliding(nums, -1)
	assert.Equal(t, [][]int{}, result)
}

func TestSlidingTwoElements(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := u.Sliding(nums, 2)
	expected := [][]int{{1, 2}, {2, 3}, {3, 4}}
	assert.Equal(t, expected, result)
}

func TestSlidingStrings(t *testing.T) {
	words := []string{"a", "b", "c", "d"}
	result := u.Sliding(words, 2)
	expected := [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}}
	assert.Equal(t, expected, result)
}

func TestSlidingDoesNotMutate(t *testing.T) {
	original := []int{1, 2, 3, 4}
	result := u.Sliding(original, 2)

	// Modify a window
	result[0][0] = 999

	// Original should be unchanged
	assert.Equal(t, 1, original[0])
}

func BenchmarkSliding(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Sliding(nums, 10)
	}
}
