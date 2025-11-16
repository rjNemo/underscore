package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestTranspose(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	result := u.Transpose(matrix)
	expected := [][]int{{1, 4}, {2, 5}, {3, 6}}
	assert.Equal(t, expected, result)
}

func TestTransposeEmpty(t *testing.T) {
	result := u.Transpose([][]int{})
	assert.Equal(t, [][]int{}, result)
}

func TestTransposeSquare(t *testing.T) {
	matrix := [][]int{{1, 2}, {3, 4}}
	result := u.Transpose(matrix)
	expected := [][]int{{1, 3}, {2, 4}}
	assert.Equal(t, expected, result)
}
