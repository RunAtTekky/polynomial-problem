package main

import (
	"encoding/json"
	"hashira/model"
	"log"
)

func main() {
	var rawData map[string]json.RawMessage
	var key model.Key
	entries := make(map[string]model.Entry)
	json_parser(&rawData, &key, &entries)
}

func json_parser(rawData *map[string]json.RawMessage, key *model.Key, entries *map[string]model.Entry) {
	input_json := `
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

	for key, val := range *rawData {
		if key == "keys" {
			continue
		}

		var value model.Entry
		err = json.Unmarshal(val, &value)

		if err != nil {
			log.Fatal("Error while parsing entry", err)
		}

		(*entries)[key] = value
	}
}
