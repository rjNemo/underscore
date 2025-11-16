package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestFoldRight(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := u.FoldRight(nums, 0, func(n, acc int) int { return n + acc })
	assert.Equal(t, 10, result)
}

func TestFoldRightEmpty(t *testing.T) {
	result := u.FoldRight([]int{}, 42, func(n, acc int) int { return n + acc })
	assert.Equal(t, 42, result)
}

func TestFoldRightSingleElement(t *testing.T) {
	result := u.FoldRight([]int{5}, 0, func(n, acc int) int { return n + acc })
	assert.Equal(t, 5, result)
}

func TestFoldRightSubtraction(t *testing.T) {
	// FoldRight: 1 - (2 - (3 - 0)) = 1 - (2 - 3) = 1 - (-1) = 2
	nums := []int{1, 2, 3}
	result := u.FoldRight(nums, 0, func(n, acc int) int { return n - acc })
	assert.Equal(t, 2, result)
}

func TestFoldRightDivision(t *testing.T) {
	// FoldRight with float: 2.0 / (4.0 / (8.0 / 1.0)) = 2.0 / (4.0 / 8.0) = 2.0 / 0.5 = 4.0
	nums := []float64{2.0, 4.0, 8.0}
	result := u.FoldRight(nums, 1.0, func(n, acc float64) float64 { return n / acc })
	assert.Equal(t, 4.0, result)
}

func TestFoldRightStrings(t *testing.T) {
	words := []string{"a", "b", "c"}
	result := u.FoldRight(words, "", func(s, acc string) string { return s + acc })
	assert.Equal(t, "abc", result)
}

func TestFoldRightVsReduce(t *testing.T) {
	nums := []int{1, 2, 3}

	// Reduce (left fold): (0 - 1) - 2 - 3 = -6
	reduceResult := u.Reduce(nums, func(n, acc int) int { return acc - n }, 0)
	assert.Equal(t, -6, reduceResult)

	// FoldRight: 1 - (2 - (3 - 0)) = 1 - (2 - 3) = 1 - (-1) = 2
	foldRightResult := u.FoldRight(nums, 0, func(n, acc int) int { return n - acc })
	assert.Equal(t, 2, foldRightResult)

	// They should be different for non-associative operations
	assert.NotEqual(t, reduceResult, foldRightResult)
}

func TestFoldRightBuildList(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.FoldRight(nums, []int{}, func(n int, acc []int) []int {
		return append([]int{n}, acc...)
	})
	assert.Equal(t, []int{1, 2, 3}, result)
}

func BenchmarkFoldRight(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.FoldRight(nums, 0, func(n, acc int) int { return n + acc })
	}
}
