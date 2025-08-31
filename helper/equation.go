package helper

import (
	"math"
)

type Equation struct {
	expression []float64
	value      float64
}

func CreateEquation(x, y, m float64) Equation {
	var equation Equation

	for i := m; i >= 0; i-- {
		x_i := math.Pow(float64(x), float64(i))

		equation.expression = append(equation.expression, x_i)
	}

	equation.value = y

	return equation
}
