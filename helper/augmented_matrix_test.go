package helper

import (
	"reflect"
	"testing"
)

func TestAugmentedMatrix(t *testing.T) {
	equations := []Equation{
		{expression: []float64{81, 27, 9, 3, 1}, value: 12},
		{expression: []float64{16, 8, 4, 2, 1}, value: 7},
	}

	got := CreateAugmentedMatrix(equations)

	want := []AugmentedEquation{
		{Expression: []float64{81, 27, 9, 3, 1, 12}},
		{Expression: []float64{16, 8, 4, 2, 1, 7}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v but want %v", got, want)
	}

}
