package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestUnique(t *testing.T) {
	nums := []int{1, 4, 2, 5, 3, 1, 5, 2, 8, 9}
	want := []int{1, 4, 2, 5, 3, 8, 9}

	assert.Equal(t, want, u.Unique(nums))
}

func TestUniqueEmpty(t *testing.T) {
	result := u.Unique([]int{})
	assert.Empty(t, result)
}

func TestUniqueSingleElement(t *testing.T) {
	result := u.Unique([]int{42})
	assert.Equal(t, []int{42}, result)
}

func TestUniqueNoDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := u.Unique(nums)
	assert.Equal(t, nums, result)
}

func TestUniqueAllSame(t *testing.T) {
	nums := []int{5, 5, 5, 5, 5}
	result := u.Unique(nums)
	assert.Equal(t, []int{5}, result)
}

func BenchmarkUnique(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i % 100 // Many duplicates
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Unique(data)
	}
}

func BenchmarkUniqueInPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := make([]int, 1000)
		for j := range data {
			data[j] = j % 100
		}
		b.StartTimer()

		u.UniqueInPlace(data)
	}
}
