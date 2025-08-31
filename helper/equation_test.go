package helper

import (
	"reflect"
	"testing"
)

func TestCreateEquation(t *testing.T) {
	equation := CreateEquation(3, 12, 4)

	want_equation := Equation{
		expression: []int{81, 27, 9, 3, 1},
		value:      12,
	}

	if !reflect.DeepEqual(equation, want_equation) {
		t.Fatalf("got %v but want %v", equation, want_equation)
	}
}
