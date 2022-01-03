package chain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rjNemo/underscore/chain"
)

func TestChainFilter(t *testing.T) {
	want := []int{2, 4, 6, 8}
	assert.Equal(t,
		want,
		chain.Of(nums).Filter(isEven).Value,
	)
}

func TestChainFilterMap(t *testing.T) {
	want := []int{4, 16, 36, 64}
	assert.Equal(t,
		want,
		chain.Of(nums).
			Filter(isEven).
			Map(toSquare).
			Value)
}

func TestChainFilterMapReduce(t *testing.T) {
	want := 120
	assert.Equal(t,
		want,
		chain.Of(nums).
			Filter(isEven).
			Map(toSquare).
			Reduce(sum, 0))
}

func TestChainFilterMapContains(t *testing.T) {
	assert.True(t, chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		Contains(16))
}

func TestChainFilterMapEach(t *testing.T) {
	want := []int{5, 17, 37, 65}
	res := make([]int, 0)
	chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		Each(func(n int) { res = append(res, n+1) })
	assert.Equal(t, want, res)
}

func TestChainFilterMapAll(t *testing.T) {
	assert.True(t, chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		All(func(n int) bool { return n%4 == 0 }))
}

func TestChainFilterMapFind(t *testing.T) {
	n, err := chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		Find(func(n int) bool { return n%4 == 0 })
	assert.Equal(t, 4, n)
	assert.NoError(t, err)
}

func TestChainFilterMapMax(t *testing.T) {
	want := 64
	assert.Equal(t, want, chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		Max())
}

func TestChainFilterMapMin(t *testing.T) {
	w := 4
	assert.Equal(t, w, chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		Min())
}

func TestChainFilterMapPartition(t *testing.T) {
	wantLeft := []int{4, 16}
	wantRight := []int{36, 64}
	left, right := chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		Partition(func(n int) bool { return n < 20 })

	assert.Equal(t, wantLeft, left)
	assert.Equal(t, wantRight, right)
}

func TestChainFilterMapAny(t *testing.T) {
	assert.True(t, chain.Of(nums).
		Filter(isEven).
		Map(toSquare).
		Any(func(n int) bool { return n%64 == 0 }))
}

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var isEven = func(n int) bool { return n%2 == 0 }
var toSquare = func(n int) int { return n * n }
var sum = func(n, acc int) int { return n + acc }
