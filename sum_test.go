package underscore_test

import (
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := 45

	assert.Equal(t, want, u.Sum(nums))
}
