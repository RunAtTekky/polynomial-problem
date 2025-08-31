package model

import (
	"encoding/json"
	"log"
)

type PolynomialInput struct {
	Key Key `json:"key"`
}

type Key struct {
	N int `json:"n"`
	K int `json:"k"`
}

type Entry struct {
	Base  string `json:"base"`
	Value string `json:"value"`
}

func Json_parser(rawData *map[string]json.RawMessage, key *Key, entries *map[string]Entry) {
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

		var value Entry
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
