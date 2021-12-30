package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestPartition(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(n int) bool { return n%2 == 0 }

	wantEvens := []int{0, 2, 4, 6, 8}
	wantOdds := []int{1, 3, 5, 7, 9}

	evens, odds := u.Partition(nums, isEven)

	assert.Equal(t, wantEvens, evens)
	assert.Equal(t, wantOdds, odds)
}
