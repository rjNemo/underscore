package underscore_test

import (
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

var zero = u.Tuple[int, string]{Left: 0, Right: "Zero"}
var one = u.Tuple[int, string]{Left: 1, Right: "One"}
var two = u.Tuple[int, string]{Left: 2, Right: "Two"}
var three = u.Tuple[int, string]{Left: 3, Right: "Three"}

func Test_Join_Can_Join_Two_Slices_Together(t *testing.T) {
	var left = []u.Tuple[int, string]{zero, one, two, three}
	var right = []u.Tuple[int, string]{one, three, two, three, two, three}

	selector := func(x u.Tuple[int, string]) int { return x.Left }

	var joined = u.Join(left, right, selector, selector)
	var want = []u.Tuple[u.Tuple[int, string], []u.Tuple[int, string]]{
		{Left: zero, Right: nil},
		{Left: one, Right: []u.Tuple[int, string]{one}},
		{Left: two, Right: []u.Tuple[int, string]{two, two}},
		{Left: three, Right: []u.Tuple[int, string]{three, three, three}},
	}

	assert.Equal(t, want, joined)
}

func Test_Join_Can_Join_and_Project_Two_Slices_Together(t *testing.T) {
	var left = []u.Tuple[int, string]{zero, one, two, three}
	var right = []u.Tuple[int, string]{one, three, two, three, two, three}

	selector := func(x u.Tuple[int, string]) int { return x.Left }
	project := func(x u.Tuple[u.Tuple[int, string], []u.Tuple[int, string]]) int {
		return len(x.Right) // projecting to a could of how many
	}

	var joined = u.JoinProject(left, right, selector, selector, project)
	var want = []int{0, 1, 2, 3}

	assert.Equal(t, want, joined)
}
