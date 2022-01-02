package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestContains(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	assert.True(t, u.Contains(nums, 5))
}

func TestNotContains(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	assert.False(t, u.Contains(nums, 15))
}
