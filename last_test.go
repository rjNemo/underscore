package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestLast(t *testing.T) {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	want := 5
	assert.Equal(t, want, u.Last(nums))
}

func TestLastEmpty(t *testing.T) {
	assert.Panics(t, func() {
		u.Last([]int{})
	})
}

func TestLastSingleElement(t *testing.T) {
	assert.Equal(t, 42, u.Last([]int{42}))
}
