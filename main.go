package main

import (
	"fmt"
	"hashira/helper"
	"hashira/model"
)

func main() {
	data := helper.TakeInput(".")
	_, key, entries := model.Json_parser(data)
	equations := helper.MakeEquations(&entries, &key)
	augmentedMatrix := helper.CreateAugmentedMatrix(equations)
	REFmatrix := helper.RowEchelonForm(augmentedMatrix)
	variables := helper.BackSubstitution(REFmatrix)

	printVariables(&variables)
}

func printVariables(variables *map[int]float64) {
	for key, val := range *variables {
		fmt.Printf("Key: %d, Value: %.3f\n", key, val)
	}
}
