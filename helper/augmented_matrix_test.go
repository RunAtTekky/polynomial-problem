package helper

import (
	"reflect"
	"testing"
)

func TestAugmentedMatrix(t *testing.T) {
	equations := []Equation{
		{expression: []int{81, 27, 9, 3, 1}, value: 12},
		{expression: []int{16, 8, 4, 1}, value: 7},
	}

	got := CreateAugmentedMatrix(equations)

	want := []AugmentedEquation{
		{Expression: []int{81, 27, 9, 3, 1, 11}},
		{Expression: []int{16, 8, 4, 1, 7}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v but want %v", got, want)
	}

}
