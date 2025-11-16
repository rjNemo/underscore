package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestTap(t *testing.T) {
	nums := []int{1, 2, 3}
	sum := 0
	result := u.Tap(nums, func(n int) { sum += n })
	assert.Equal(t, nums, result)
	assert.Equal(t, 6, sum)
}

func TestTapEmpty(t *testing.T) {
	result := u.Tap([]int{}, func(n int) {})
	assert.Equal(t, []int{}, result)
}
