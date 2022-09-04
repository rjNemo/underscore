package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func Test_OrderBy_Asc(t *testing.T) {
	set := u.Range(5, 0)
	want := u.Range(0, 5)

	result := u.OrderBy(set, func(left int, right int) bool {
		return left > right
	})

	assert.Equal(t, want, result)
}

func Test_OrderBy_Desc(t *testing.T) {
	set := u.Range(0, 5)
	want := u.Range(5, 0)

	result := u.OrderBy(set, func(left int, right int) bool {
		return left < right
	})

	assert.Equal(t, want, result)
}
