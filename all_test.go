package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestAll(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	isOdd := func(n int) bool { return n%2 != 0 }
	assert.True(t, u.All(nums, isOdd))
}

func TestNotAll(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 10}
	isOdd := func(n int) bool { return n%2 != 0 }
	assert.False(t, u.All(nums, isOdd))
}
