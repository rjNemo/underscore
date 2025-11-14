package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestDrop(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	want := []int{3, 4, 5}

	assert.Equal(t, want, u.Drop(nums, 2))
}

func TestDropNone(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	assert.Equal(t, nums, u.Drop(nums, 0))
	assert.Equal(t, nums, u.Drop(nums, -1))
}

func TestDropAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	assert.Empty(t, u.Drop(nums, 5))
	assert.Empty(t, u.Drop(nums, 10))
}

func TestDropEmpty(t *testing.T) {
	result := u.Drop([]int{}, 5)
	assert.Empty(t, result)
}

func TestDropSingleElement(t *testing.T) {
	nums := []int{42}

	assert.Equal(t, nums, u.Drop(nums, 0))
	assert.Empty(t, u.Drop(nums, 1))
}
