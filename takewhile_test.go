package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestTakeWhile(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := u.TakeWhile(nums, func(n int) bool { return n < 5 })
	assert.Equal(t, []int{1, 2, 3, 4}, result)
}

func TestTakeWhileEmpty(t *testing.T) {
	result := u.TakeWhile([]int{}, func(n int) bool { return n < 5 })
	assert.Equal(t, []int{}, result)
}

func TestTakeWhileNoneMatch(t *testing.T) {
	nums := []int{5, 6, 7, 8, 9}
	result := u.TakeWhile(nums, func(n int) bool { return n < 5 })
	assert.Equal(t, []int{}, result)
}

func TestTakeWhileAllMatch(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := u.TakeWhile(nums, func(n int) bool { return n < 10 })
	assert.Equal(t, []int{1, 2, 3, 4}, result)
}

func TestTakeWhileSingleElement(t *testing.T) {
	result := u.TakeWhile([]int{5}, func(n int) bool { return n < 10 })
	assert.Equal(t, []int{5}, result)
}

func TestTakeWhileStrings(t *testing.T) {
	words := []string{"apple", "banana", "cherry", "date"}
	result := u.TakeWhile(words, func(s string) bool { return len(s) < 6 })
	assert.Equal(t, []string{"apple"}, result)
}

func BenchmarkTakeWhile(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.TakeWhile(nums, func(n int) bool { return n < 500 })
	}
}
