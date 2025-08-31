package helper

import (
	"hashira/model"
	"math"
	"sort"
	"strconv"
)

type Equation struct {
	expression []float64
	value      float64
}

type byColumn [][]float64

func (a byColumn) Len() int {
	return len(a)
}

func (a byColumn) Less(i, j int) bool {
	return a[i][len(a[i])-1] < a[j][len(a[j])-1]
}

func (a byColumn) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
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

	var allEntries byColumn
	for x, entry := range *entries {
		y := BaseConversion(entry.Base, entry.Value)
		xInt, _ := strconv.Atoi(x)

		allEntries = append(allEntries, []float64{float64(xInt), y})
	}

	sort.Sort(allEntries)

	for _, entry := range allEntries {
		x := entry[0]
		y := entry[1]

		equation := CreateEquation(x, y, float64(key.K-1))
		equations = append(equations, equation)

		if len(equations) >= key.K {
			break
		}
	}

	return equations
}
