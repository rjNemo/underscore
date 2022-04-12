package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestUnique(t *testing.T) {
	nums := []int{1, 4, 2, 5, 3, 1, 5, 2, 8, 9}
	want := []int{1, 4, 2, 5, 3, 8, 9}

	assert.Equal(t, want, u.Unique(nums))
}
