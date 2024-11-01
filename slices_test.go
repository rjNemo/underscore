package underscore_test

import (
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

func TestSortSliceAscString(t *testing.T) {
	slc := []string{"c", "a", "b"}
	expected := []string{"a", "b", "c"}
	u.SortSliceASC(slc)

	assert.Equal(t, expected, slc)
}

func TestSortSliceDescString(t *testing.T) {
	slc := []string{"c", "a", "b"}
	expected := []string{"c", "b", "a"}
	u.SortSliceDESC(slc)

	assert.Equal(t, expected, slc)
}

func TestSortSliceAscInt(t *testing.T) {
	slc := []int{1, 4, 3, 5, 2}
	expected := []int{1, 2, 3, 4, 5}
	u.SortSliceASC(slc)

	assert.Equal(t, expected, slc)
}

func TestSortSliceDescInt(t *testing.T) {
	slc := []int{1, 4, 3, 5, 2}
	expected := []int{5, 4, 3, 2, 1}
	u.SortSliceDESC(slc)

	assert.Equal(t, expected, slc)
}

func TestSortSliceAscFloat64(t *testing.T) {
	slc := []float64{1.0, 1.2, 1.1, 1.5, 1.01}
	expected := []float64{1, 1.01, 1.1, 1.2, 1.5}
	u.SortSliceASC(slc)

	assert.Equal(t, expected, slc)
}

func TestSortSliceDescFloat64(t *testing.T) {
	slc := []float64{1.0, 1.2, 1.1, 1.5, 1.01}
	expected := []float64{1.5, 1.2, 1.1, 1.01, 1}
	u.SortSliceDESC(slc)

	assert.Equal(t, expected, slc)
}
