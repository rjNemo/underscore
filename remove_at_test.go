package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestRemoveAt(t *testing.T) {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	want := []int{1, 9, 2, 3, 7, 4, 6, 5}

	assert.Equal(t, want, u.RemoveAt(nums, 3))
}

func TestRemoveAtFirst(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	want := []int{2, 3, 4, 5}

	assert.Equal(t, want, u.RemoveAt(nums, 0))
}

func TestRemoveAtLast(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4}

	assert.Equal(t, want, u.RemoveAt(nums, 4))
}

func TestRemoveAtOutOfBounds(t *testing.T) {
	nums := []int{1, 2, 3}

	// Negative index
	assert.Equal(t, nums, u.RemoveAt(nums, -1))

	// Index too large
	assert.Equal(t, nums, u.RemoveAt(nums, 10))
}

func TestRemoveAtEmpty(t *testing.T) {
	result := u.RemoveAt([]int{}, 0)
	assert.Empty(t, result)
}

func TestRemoveAtSingleElement(t *testing.T) {
	result := u.RemoveAt([]int{42}, 0)
	assert.Empty(t, result)
}
