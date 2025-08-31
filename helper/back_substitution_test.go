package helper

import (
	"math"
	"testing"
)

func TestBackSubstitution(t *testing.T) {
	REFMatrix := []AugmentedEquation{
		{Expression: []float64{36, 6, 1, 39}},
		{Expression: []float64{0, 1.5, 0.75, 2.25}},
		{Expression: []float64{0, 0, 0.22222, 0.66667}},
	}

	got := BackSubstitution(REFMatrix)

	want := map[int]float64{
		0: 1,
		1: 0,
		2: 3,
	}

	if !AreMapEqual(got, want) {
		t.Errorf("got %v\nbut want %v", got, want)
	}
}

func AreMapEqual(mapA, mapB map[int]float64) bool {
	if len(mapA) != len(mapB) {
		return false
	}

	for key, val := range mapA {
		valB := mapB[key]

		threshold := 1e-3

		if math.Abs(val-valB) > threshold {
			return false
		}
	}

	return true
}
