package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestDifference(t *testing.T) {
	nums := []int{1, 3, 5, 6, 7, 9}
	reject := []int{9, 7, 5, 4}
	want := []int{1, 3, 6}

	assert.Equal(t, want, u.Difference(nums, reject))
}
