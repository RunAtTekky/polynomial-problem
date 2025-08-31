package main

import (
	"encoding/json"
	"fmt"
	"hashira/helper"
	"hashira/model"
	"log"
	"strconv"
)

func main() {
	var rawData map[string]json.RawMessage
	var key model.Key
	entries := make(map[string]model.Entry)
	json_parser(&rawData, &key, &entries)
	equations := makeEquations(&entries, &key)
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

func makeEquations(entries *map[string]model.Entry, key *model.Key) []helper.Equation {
	var equations []helper.Equation

	for x, entry := range *entries {
		y := helper.BaseConversion(entry.Base, entry.Value)

		xInt, _ := strconv.Atoi(x)

		equation := helper.CreateEquation(float64(xInt), y, float64(key.K-1))
		equations = append(equations, equation)
	}

	return equations
}

func json_parser(rawData *map[string]json.RawMessage, key *model.Key, entries *map[string]model.Entry) {
	input_json := `
	{
	"keys": {
        "n": 4,
        "k": 3
    },
    "1": {
        "base": "10",
        "value": "4"
    },
    "2": {
        "base": "2",
        "value": "111"
    },
    "3": {
        "base": "10",
        "value": "12"
    },
    "6": {
        "base": "4",
        "value": "213"
    }
}
	`

	err := json.Unmarshal([]byte(input_json), &rawData)
	if err != nil {
		log.Fatal("Error while parsing json file", err)
	}

	err = json.Unmarshal((*rawData)["keys"], &key)
	if err != nil {
		log.Fatal("Error while parsing keys", err)
	}

	for k, val := range *rawData {
		if k == "keys" {
			continue
		}

		var value model.Entry
		err = json.Unmarshal(val, &value)

		if err != nil {
			log.Fatal("Error while parsing entry", err)
		}

		(*entries)[k] = value

		if len(*entries) >= key.K {
			break
		}
	}
}
