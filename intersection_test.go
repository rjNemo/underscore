package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestIntersection(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 3, 5, 8, 0}
	want := []int{3, 5}

	assert.Equal(t, want, u.Intersection(a, b))
}
