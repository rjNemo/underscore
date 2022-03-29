package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipingExample(t *testing.T) {
	assert.Equal(t, 120, piping())
}
