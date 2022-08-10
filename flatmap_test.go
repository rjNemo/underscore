package underscore_test

import (
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

func TestFlatmap(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	transform := func(n int) []int { return []int{(n - 1) * n, (n) * n} }
	want := []int{0, 1, 2, 4, 6, 9, 12, 16}

	assert.Equal(t, want, u.Flatmap(nums, transform))
}
