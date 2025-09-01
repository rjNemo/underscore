package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestUniqueInPlace(t *testing.T) {
	nums := []int{1, 4, 2, 5, 3, 1, 5, 2, 8, 9}
	got := u.UniqueInPlace(nums)
	want := []int{1, 4, 2, 5, 3, 8, 9}
	assert.Equal(t, want, got)
}
