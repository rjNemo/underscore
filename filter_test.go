package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestFilter(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(n int) bool { return n%2 == 0 }

	want := []int{0, 2, 4, 6, 8}
	assert.Equal(t, want, u.Filter(nums, isEven))
}

func BenchmarkFilter(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	isEven := func(n int) bool { return n%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Filter(data, isEven)
	}
}
