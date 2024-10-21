package underscore_test

import (
	"reflect"
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

func TestPointers(t *testing.T) {
	variable := 123
	var object struct{}

	cases := []struct {
		value    any
		expected bool
	}{
		{
			value:    u.ToPointer("myValue"),
			expected: true,
		},
		{
			value:    u.ToPointer(variable),
			expected: true,
		},
		{
			value:    &variable,
			expected: true,
		},
		{
			value:    nil,
			expected: false,
		},
		{
			value:    u.ToPointer(object),
			expected: true,
		},
	}

	for _, c := range cases {
		got := (reflect.ValueOf(c.value).Kind() == reflect.Ptr)
		assert.Equal(t, c.expected, got)
	}
}
