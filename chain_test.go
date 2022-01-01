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

	assert.True(t, u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Contains(16))

	want = []int{5, 17, 37, 65}
	res := make([]int, 0)
	u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Each(func(n int) {
			res = append(res, n+1)
		})
	assert.Equal(t, want, res)

	assert.True(t, u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Every(func(n int) bool { return n%4 == 0 }))

	f, err := u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Find(func(n int) bool { return n%4 == 0 })
	assert.Equal(t, 4, f)
	assert.NoError(t, err)

	w = 64
	assert.Equal(t, w, u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Max())

	w = 4
	assert.Equal(t, w, u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Min())

	wFirstHalf := []int{4, 16}
	wSecondHalf := []int{36, 64}
	firstHalf, secondHalf := u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Partition(func(n int) bool { return n < 20 })

	assert.Equal(t, wFirstHalf, firstHalf)
	assert.Equal(t, wSecondHalf, secondHalf)

	assert.True(t, u.NewChain(nums).
		Filter(isEven).
		Map(toSquare).
		Some(func(n int) bool { return n%64 == 0 }))
}
