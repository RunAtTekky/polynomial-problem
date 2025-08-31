package helper

import (
	"math"
)

type Equation struct {
	expression []int
	value      int
}

func CreateEquation(x, y, m int) Equation {
	var equation Equation

	for i := m; i >= 0; i-- {
		x_i := math.Pow(float64(x), float64(i))

		equation.expression = append(equation.expression, int(x_i))
	}

	equation.value = y

	return equation
}
