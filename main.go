package main

import (
	"hashira/helper"
	"hashira/model"
)

func main() {
	PolynomialFinder()
}

func PolynomialFinder() {
	data := helper.TakeInput(".")
	_, key, entries := model.Json_parser(data)
	equations := helper.MakeEquations(&entries, &key)
	augmentedMatrix := helper.CreateAugmentedMatrix(equations)
	REFmatrix := helper.RowEchelonForm(augmentedMatrix)
	variables := helper.BackSubstitution(REFmatrix)

	helper.PrintEquation(key.K-1, &variables)
}
