package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestMin(t *testing.T) {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	want := 1
	assert.Equal(t, want, u.Min(nums))
}
