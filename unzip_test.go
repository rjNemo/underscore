package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestUnzip(t *testing.T) {
	pairs := []u.Tuple[int, string]{
		{Left: 1, Right: "a"},
		{Left: 2, Right: "b"},
		{Left: 3, Right: "c"},
	}
	lefts, rights := u.Unzip(pairs)
	assert.Equal(t, []int{1, 2, 3}, lefts)
	assert.Equal(t, []string{"a", "b", "c"}, rights)
}

func TestUnzipEmpty(t *testing.T) {
	lefts, rights := u.Unzip([]u.Tuple[int, string]{})
	assert.Equal(t, []int{}, lefts)
	assert.Equal(t, []string{}, rights)
}
