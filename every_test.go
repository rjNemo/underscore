package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestEvery(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	isOdd := func(n int) bool { return n%2 != 0 }
	assert.True(t, u.Every(nums, isOdd))
}
