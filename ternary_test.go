package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestTernary(t *testing.T) {
	tests := []struct {
		condition bool
		want      string
	}{
		{true, "foo"},
		{false, "bar"},
	}

	for _, tc := range tests {
		assert.Equal(t, u.Ternary(tc.condition, "foo", "bar"), tc.want)
	}
}
