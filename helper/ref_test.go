package helper

import (
	"math"
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

	if !areMatricesEqual(got, want) {
		t.Fatalf("got \n%v\n but want \n%v\n given \n%v\n", got, want, augmentedMatrix)
	}
}

func areMatricesEqual(matA, matB []AugmentedEquation) bool {
	if len(matA) != len(matB) {
		return false
	}

	if len(matA[0].Expression) != len(matB[0].Expression) {
		return false
	}

	threshold := 1e-3

	for row := range matA {
		for col := range matA[0].Expression {
			if math.Abs(matA[row].Expression[col]-matB[row].Expression[col]) > threshold {
				return false
			}
		}
	}

	return true
}
