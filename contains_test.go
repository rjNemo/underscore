package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestContains(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	assert.True(t, u.Contains(nums, 5))
}

func TestNotContains(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	assert.False(t, u.Contains(nums, 15))
}

func TestContainsBy(t *testing.T) {
	nums := []int{1, 3, 5, 8}
	assert.True(t, u.ContainsBy(nums, func(n int) bool { return n%2 == 0 }))
	assert.False(t, u.ContainsBy(nums, func(n int) bool { return n < 0 }))
}

func TestContainsByStruct(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	users := []user{{1, "a"}, {2, "b"}, {3, "c"}}
	assert.True(t, u.ContainsBy(users, func(u user) bool { return u.ID == 2 }))
	assert.False(t, u.ContainsBy(users, func(u user) bool { return u.Name == "z" }))
}
