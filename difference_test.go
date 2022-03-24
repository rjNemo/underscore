package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestDifference(t *testing.T) {
	a := []int{1, 3, 5, 6, 7, 9}
	b := []int{9, 7, 5, 4}
	want := []int{1, 3, 6}

	assert.Equal(t, want, u.Difference(a, b))
}
