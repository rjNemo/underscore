package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestReplicate(t *testing.T) {
	result := u.Replicate(3, "hello")
	assert.Equal(t, []string{"hello", "hello", "hello"}, result)
}

func TestReplicateZero(t *testing.T) {
	result := u.Replicate(0, 42)
	assert.Equal(t, []int{}, result)
}

func TestReplicateNegative(t *testing.T) {
	result := u.Replicate(-5, 42)
	assert.Equal(t, []int{}, result)
}

func TestReplicateOne(t *testing.T) {
	result := u.Replicate(1, 100)
	assert.Equal(t, []int{100}, result)
}
