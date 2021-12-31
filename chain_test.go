package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestChain(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(n int) bool { return n%2 == 0 }
	toSquare := func(n int) int { return n * n }
	sum := func(n, acc int) int { return n + acc }

	want := []int{2, 4, 6, 8}
	assert.Equal(t, want, u.NewChain(nums).
		Filter(isEven).
		Value)

	want = []int{4, 16, 36, 64}
	assert.Equal(t, want, u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Value)

	w := 120
	assert.Equal(t, w, u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Reduce(sum, 0))
}
