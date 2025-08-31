package helper

import (
	"reflect"
	"testing"
)

func TestRowEchelonForm(t *testing.T) {

	augmentedMatrix := []AugmentedEquation{
		{Expression: []float64{36, 6, 1, 39}},
		{Expression: []float64{9, 3, 1, 12}},
		{Expression: []float64{4, 2, 1, 7}},
		// {Expression: []float64{1, 1, 1, 4}},
	}

	want := []AugmentedEquation{
		{Expression: []float64{36, 6, 1, 39}},
		{Expression: []float64{0, 1.5, 0.75, 2.25}},
		{Expression: []float64{0, 0, 0.22222, 0.66667}},
		// {Expression: []float64{0, 0, 0, 0}},
	}

	got := RowEchelonForm(augmentedMatrix)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got \n%v\n but want \n%v\n given \n%v\n", got, want, augmentedMatrix)
	}

}
