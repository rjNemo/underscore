package underscore_test

import (
	"fmt"
	"testing"

	"github.com/rjNemo/underscore"

	"github.com/stretchr/testify/assert"
)

func TestEach(t *testing.T) {
	names := []string{"Alice", "Bob", "Charles"}
	want := []string{"Hi Alice", "Hi Bob", "Hi Charles"}

	res := make([]string, 0)
	underscore.Each(names, func(n string) {
		res = append(res, fmt.Sprintf("Hi %s", n))
	})

	assert.Equal(t, want, res)
}

func TestEachReturnsInitialSlice(t *testing.T) {
	names := []string{"Alice", "Bob", "Charles"}
	want := []string{"Alice", "Bob", "Charles"}

	res := make([]string, 0)

	assert.Equal(t, want, underscore.Each(names, func(n string) {
		res = append(res, fmt.Sprintf("Hi %s", n))
	}))
}
