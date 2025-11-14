package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func Test_OrderBy_Asc(t *testing.T) {
	set := u.Range(5, 0)
	want := u.Range(0, 5)

	result := u.OrderBy(set, func(left int, right int) bool {
		return left > right
	})

	assert.Equal(t, want, result)
}

func Test_OrderBy_Desc(t *testing.T) {
	set := u.Range(0, 5)
	want := u.Range(5, 0)

	result := u.OrderBy(set, func(left int, right int) bool {
		return left < right
	})

	assert.Equal(t, want, result)
}

func BenchmarkOrderBy(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = 1000 - i // Reverse order - worst case for bubble sort
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dataCopy := make([]int, len(data))
		copy(dataCopy, data)
		u.OrderBy(dataCopy, func(a, b int) bool { return a > b })
	}
}

func BenchmarkOrderBySmall(b *testing.B) {
	data := make([]int, 10)
	for i := range data {
		data[i] = 10 - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dataCopy := make([]int, len(data))
		copy(dataCopy, data)
		u.OrderBy(dataCopy, func(a, b int) bool { return a > b })
	}
}
