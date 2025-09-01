package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	m "github.com/rjNemo/underscore/maps"
)

func TestKeysValues(t *testing.T) {
	in := map[int]string{1: "a", 2: "b", 3: "c"}
	ks := m.Keys(in)
	vs := m.Values(in)

	// Order is unspecified; verify content and lengths.
	assert.Len(t, ks, 3)
	assert.ElementsMatch(t, []int{1, 2, 3}, ks)

	assert.Len(t, vs, 3)
	assert.ElementsMatch(t, []string{"a", "b", "c"}, vs)
}
