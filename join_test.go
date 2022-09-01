package underscore_test

import (
	"reflect"
	"testing"

	u "github.com/rjNemo/underscore"
)

func Test_Join_Can_Join_Two_Slices_Together(t *testing.T) {
	one := u.Tuple[int, string]{Left: 1, Right: "One"}
	two := u.Tuple[int, string]{Left: 2, Right: "Two"}
	three := u.Tuple[int, string]{Left: 3, Right: "Three"}

	var left = []u.Tuple[int, string]{one, two, three}
	var right = []u.Tuple[int, string]{one, three, two, three, two, three}

	selector := func(x u.Tuple[int, string]) int { return x.Left }

	var joined = u.Join(left, right, selector, selector)
	var want = []u.Tuple[u.Tuple[int, string], []u.Tuple[int, string]]{
		{Left: one, Right: []u.Tuple[int, string]{one}},
		{Left: two, Right: []u.Tuple[int, string]{two, two}},
		{Left: three, Right: []u.Tuple[int, string]{three, three, three}},
	}

	if !reflect.DeepEqual(joined, want) {
		t.Errorf("Expected to get %v but we got %v instead", want, joined)
	}
}
