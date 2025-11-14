package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestInit(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	init, last := u.Init(nums)
	assert.Equal(t, []int{1, 2, 3, 4}, init)
	assert.Equal(t, 5, last)
}

func TestInitEmpty(t *testing.T) {
	init, last := u.Init([]int{})
	assert.Equal(t, []int{}, init)
	assert.Equal(t, 0, last)
}

func TestInitSingleElement(t *testing.T) {
	init, last := u.Init([]int{42})
	assert.Equal(t, []int{}, init)
	assert.Equal(t, 42, last)
}

func TestInitTwoElements(t *testing.T) {
	init, last := u.Init([]int{1, 2})
	assert.Equal(t, []int{1}, init)
	assert.Equal(t, 2, last)
}

func TestInitStrings(t *testing.T) {
	words := []string{"hello", "world", "!"}
	init, last := u.Init(words)
	assert.Equal(t, []string{"hello", "world"}, init)
	assert.Equal(t, "!", last)
}

func TestInitDoesNotMutate(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	init, last := u.Init(original)

	// Modify returned slice
	init[0] = 999

	// Original should be unchanged
	assert.Equal(t, 1, original[0])
	assert.Equal(t, 5, last)
}

func BenchmarkInit(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Init(nums)
	}
}
