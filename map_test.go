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
