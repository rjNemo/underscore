package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMapReduceExample(t *testing.T) {
	assert.Equal(t, 120, filterMapReduce())
}
