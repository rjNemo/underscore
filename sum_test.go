package underscore_test

import (
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := 45

	assert.Equal(t, want, u.Sum(nums))
}

func TestSumMap(t *testing.T) {
	nums := []u.Tuple[string, int]{
		{"zero", 0},
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}
	want := 45

	assert.Equal(t, want, u.SumMap(nums, func(item u.Tuple[string, int]) int { return item.Right }))
}
