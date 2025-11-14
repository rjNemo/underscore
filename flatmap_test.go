package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestFlatmap(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	transform := func(n int) []int { return []int{(n - 1) * n, (n) * n} }
	want := []int{0, 1, 2, 4, 6, 9, 12, 16}

	assert.Equal(t, want, u.Flatmap(nums, transform))
}

func BenchmarkFlatmap(b *testing.B) {
	data := make([]int, 100)
	for i := range data {
		data[i] = i
	}
	mapper := func(n int) []int { return []int{n, n * 2, n * 3} }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Flatmap(data, mapper)
	}
}
