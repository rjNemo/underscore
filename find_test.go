package underscore_test

import (
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	want := 5
	nums := []int{2, 4, 5, 6, 8, 0}
	isOdd := func(n int) bool { return n%2 != 0 }

	got, err := u.Find(nums, isOdd)
	assert.Equal(t, want, got)
	assert.NoError(t, err)
}

func TestNotFound(t *testing.T) {
	nums := []int{2, 4, 6, 8, 0}
	isOdd := func(n int) bool { return n%2 != 0 }

	_, err := u.Find(nums, isOdd)
	assert.Error(t, err)
}
