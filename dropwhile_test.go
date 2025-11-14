package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestDropWhile(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := u.DropWhile(nums, func(n int) bool { return n < 5 })
	assert.Equal(t, []int{5, 6, 7, 8, 9}, result)
}

func TestDropWhileEmpty(t *testing.T) {
	result := u.DropWhile([]int{}, func(n int) bool { return n < 5 })
	assert.Equal(t, []int{}, result)
}

func TestDropWhileNoneMatch(t *testing.T) {
	nums := []int{5, 6, 7, 8, 9}
	result := u.DropWhile(nums, func(n int) bool { return n < 5 })
	assert.Equal(t, []int{5, 6, 7, 8, 9}, result)
}

func TestDropWhileAllMatch(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := u.DropWhile(nums, func(n int) bool { return n < 10 })
	assert.Equal(t, []int{}, result)
}

func TestDropWhileSingleElement(t *testing.T) {
	result := u.DropWhile([]int{5}, func(n int) bool { return n < 10 })
	assert.Equal(t, []int{}, result)
}

func TestDropWhileStrings(t *testing.T) {
	words := []string{"apple", "banana", "cherry", "date"}
	result := u.DropWhile(words, func(s string) bool { return len(s) < 6 })
	assert.Equal(t, []string{"banana", "cherry", "date"}, result)
}

func BenchmarkDropWhile(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.DropWhile(nums, func(n int) bool { return n < 500 })
	}
}
