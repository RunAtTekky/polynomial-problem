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

func Json_parser(input_json string) (map[string]json.RawMessage, Key, map[string]Entry) {
	var rawData map[string]json.RawMessage
	err := json.Unmarshal([]byte(input_json), &rawData)
	if err != nil {
		log.Fatal("Error while parsing json file", err)
	}

	var key Key
	err = json.Unmarshal(rawData["keys"], &key)
	if err != nil {
		log.Fatal("Error while parsing keys", err)
	}

	entries := make(map[string]Entry)
	for k, val := range rawData {
		if k == "keys" {
			continue
		}

		var value Entry
		err = json.Unmarshal(val, &value)

		if err != nil {
			log.Fatal("Error while parsing entry", err)
		}

		entries[k] = value

		if len(entries) >= key.K {
			break
		}
	}

	return rawData, key, entries
}
