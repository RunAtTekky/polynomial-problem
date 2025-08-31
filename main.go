package main

import (
	"encoding/json"
	"fmt"
	"hashira/helper"
	"hashira/model"
)

func main() {
	var rawData map[string]json.RawMessage
	var key model.Key
	entries := make(map[string]model.Entry)
	model.Json_parser(&rawData, &key, &entries)
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
