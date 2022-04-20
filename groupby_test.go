package underscore_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestGroupBy(t *testing.T) {
	nums := []float64{1.3, 2.1, 2.4}
	want := map[int][]float64{
		1: {1.3},
		2: {2.1, 2.4},
	}
	f := func(n float64) int {
		return int(math.Floor(n))
	}
	assert.Equal(t, want, u.GroupBy(nums, f))
}
