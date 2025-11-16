package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestIntersperse(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := u.Intersperse(nums, 0)
	assert.Equal(t, []int{1, 0, 2, 0, 3, 0, 4, 0, 5}, result)
}

func TestIntersperseEmpty(t *testing.T) {
	result := u.Intersperse([]int{}, 0)
	assert.Equal(t, []int{}, result)
}

func TestIntersperseSingleElement(t *testing.T) {
	result := u.Intersperse([]int{42}, 0)
	assert.Equal(t, []int{42}, result)
}

func TestIntersperseTwoElements(t *testing.T) {
	result := u.Intersperse([]int{1, 2}, 0)
	assert.Equal(t, []int{1, 0, 2}, result)
}

func TestIntersperseStrings(t *testing.T) {
	words := []string{"hello", "world", "!"}
	result := u.Intersperse(words, ",")
	assert.Equal(t, []string{"hello", ",", "world", ",", "!"}, result)
}

func TestIntersperseComma(t *testing.T) {
	words := []string{"apple", "banana", "cherry"}
	result := u.Intersperse(words, ",")
	assert.Equal(t, []string{"apple", ",", "banana", ",", "cherry"}, result)
}

func TestIntersperseNegativeNumber(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.Intersperse(nums, -1)
	assert.Equal(t, []int{1, -1, 2, -1, 3}, result)
}

func BenchmarkIntersperse(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Intersperse(nums, 0)
	}
}
