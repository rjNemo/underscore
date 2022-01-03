package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainingExample(t *testing.T) {
	assert.Equal(t, 120, chaining())
}
