package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestChunk(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	got := u.Chunk(in, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	assert.Equal(t, want, got)
}

func TestChunkLargeSize(t *testing.T) {
	in := []int{1, 2, 3}
	got := u.Chunk(in, 10)
	want := [][]int{{1, 2, 3}}
	assert.Equal(t, want, got)
}

func TestChunkInvalidSize(t *testing.T) {
	var in []int
	assert.Nil(t, u.Chunk(in, 0))
	assert.Nil(t, u.Chunk(in, -1))
}

func TestChunkEmpty(t *testing.T) {
	got := u.Chunk([]int{}, 1)
	assert.Equal(t, 0, len(got))
}
