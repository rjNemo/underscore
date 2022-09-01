package underscore_test

import (
	"reflect"
	"testing"

	u "github.com/rjNemo/underscore"
)

func Test_Range_Creates_Slices(t *testing.T) {
	range1 := u.Range(0, 5)
	want1 := []int{0, 1, 2, 3, 4, 5}

	if !reflect.DeepEqual(range1, want1) {
		t.Errorf("Expected the result to be %v but we got %v", want1, range1)
	}

	range2 := u.Range(5, 0)
	want2 := []int{5, 4, 3, 2, 1, 0}

	if !reflect.DeepEqual(range2, want2) {
		t.Errorf("Expected the result to be %v but we got %v", want2, range2)
	}
}
