package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestUniqueBy(t *testing.T) {
	type user struct {
		ID    int
		Email string
	}
	in := []user{{1, "a@x"}, {2, "b@x"}, {3, "a@x"}, {4, "c@x"}, {5, "b@x"}}
	out := u.UniqueBy(in, func(u user) string { return u.Email })
	want := []user{{1, "a@x"}, {2, "b@x"}, {4, "c@x"}}
	assert.Equal(t, want, out)
}
