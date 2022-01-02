package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestSome(t *testing.T) {
	nums := []int{1, 2, 4, 6, 8}
	isOdd := func(n int) bool { return n%2 != 0 }

	assert.True(t, u.Some(nums, isOdd))
}

func TestNotSome(t *testing.T) {
	nums := []int{2, 4, 6, 8}
	isOdd := func(n int) bool { return n%2 != 0 }

	assert.False(t, u.Some(nums, isOdd))
}
