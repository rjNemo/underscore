package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3}
	f := func(n int) int {
		return n * n
	}
	want := []int{1, 4, 9}
	assert.Equal(t, want, u.Map(nums, f))
}

func TestMapEmpty(t *testing.T) {
	result := u.Map([]int{}, func(n int) int { return n * 2 })
	assert.Empty(t, result)
}

func TestMapSingleElement(t *testing.T) {
	result := u.Map([]int{5}, func(n int) int { return n * 2 })
	assert.Equal(t, []int{10}, result)
}

func TestMapLarge(t *testing.T) {
	large := make([]int, 10000)
	for i := range large {
		large[i] = i
	}
	result := u.Map(large, func(n int) int { return n * 2 })
	assert.Equal(t, 10000, len(result))
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 19998, result[9999])
}
