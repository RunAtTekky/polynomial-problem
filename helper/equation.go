package helper

import (
	"hashira/model"
	"math"
	"strconv"
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

func MakeEquations(entries *map[string]model.Entry, key *model.Key) []Equation {
	var equations []Equation

	for x, entry := range *entries {
		y := BaseConversion(entry.Base, entry.Value)

		xInt, _ := strconv.Atoi(x)

		equation := CreateEquation(float64(xInt), y, float64(key.K-1))
		equations = append(equations, equation)
	}

	return equations
}
