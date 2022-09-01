package underscore_test

import (
	"reflect"
	"testing"

	u "github.com/rjNemo/underscore"
)

func Test_Zip_Can_Zip_Two_Equal_Sized_Slices(t *testing.T) {
	left := []string{"Left 1", "Left 2", "Left 3"}
	right := []int{1, 2, 3}

	var zipped = u.Zip(left, right)

	want := []u.Tuple[string, int]{
		{Left: "Left 1", Right: 1},
		{Left: "Left 2", Right: 2},
		{Left: "Left 3", Right: 3},
	}

	if !reflect.DeepEqual(zipped, want) {
		t.Errorf("Expected the result to be %v but we got %v", want, zipped)
	}
}

func Test_Zip_Can_Zip_Two_Different_Sized_Slices_Left_Larger(t *testing.T) {
	left := []string{"Left 1", "Left 2", "Left 3", "Left 4"}
	right := []int{1, 2, 3}

	var zipped = u.Zip(left, right)
	if len(zipped) != 3 {
		t.Errorf("Expected the result of Zip(left, right) to have a length of 3 but got %v", len(zipped))
	}

	want := []u.Tuple[string, int]{
		{Left: "Left 1", Right: 1},
		{Left: "Left 2", Right: 2},
		{Left: "Left 3", Right: 3},
	}

	if !reflect.DeepEqual(zipped, want) {
		t.Errorf("Expected the result to be %v but we got %v", want, zipped)
	}
}

func Test_Zip_Can_Zip_Two_Different_Sized_Slices_Right_Larger(t *testing.T) {
	left := []string{"Left 1", "Left 2", "Left 3"}
	right := []int{1, 2, 3, 4}

	var zipped = u.Zip(left, right)
	if len(zipped) != 3 {
		t.Errorf("Expected the result of Zip(left, right) to have a length of 3 but got %v", len(zipped))
	}

	want := []u.Tuple[string, int]{
		{Left: "Left 1", Right: 1},
		{Left: "Left 2", Right: 2},
		{Left: "Left 3", Right: 3},
	}

	if !reflect.DeepEqual(zipped, want) {
		t.Errorf("Expected the result to be %v but we got %v", want, zipped)
	}
}
