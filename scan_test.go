package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestScan(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := u.Scan(nums, 0, func(acc, n int) int { return acc + n })
	assert.Equal(t, []int{1, 3, 6, 10}, result)
}

func TestScanEmpty(t *testing.T) {
	result := u.Scan([]int{}, 0, func(acc, n int) int { return acc + n })
	assert.Equal(t, []int{}, result)
}

func TestScanSingleElement(t *testing.T) {
	result := u.Scan([]int{5}, 0, func(acc, n int) int { return acc + n })
	assert.Equal(t, []int{5}, result)
}

func TestScanMultiplication(t *testing.T) {
	nums := []int{2, 3, 4}
	result := u.Scan(nums, 1, func(acc, n int) int { return acc * n })
	assert.Equal(t, []int{2, 6, 24}, result)
}

func TestScanStrings(t *testing.T) {
	words := []string{"hello", "world", "!"}
	result := u.Scan(words, "", func(acc, s string) string { return acc + s })
	assert.Equal(t, []string{"hello", "helloworld", "helloworld!"}, result)
}

func TestScanMax(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5, 9, 2}
	result := u.Scan(nums, 0, func(acc, n int) int {
		if n > acc {
			return n
		}
		return acc
	})
	assert.Equal(t, []int{3, 3, 4, 4, 5, 9, 9}, result)
}

func TestScanDifferentTypes(t *testing.T) {
	nums := []int{1, 2, 3}
	result := u.Scan(nums, 0.0, func(acc float64, n int) float64 {
		return acc + float64(n)*2.5
	})
	assert.Equal(t, []float64{2.5, 7.5, 15.0}, result)
}

func BenchmarkScan(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Scan(nums, 0, func(acc, n int) int { return acc + n })
	}
}
