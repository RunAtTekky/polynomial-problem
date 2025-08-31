package helper

import (
	"hashira/model"
	"reflect"
	"testing"
)

func TestInput(t *testing.T) {
	want := `{
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
}`
	_, wantKey, wantEntries := model.Json_parser(want)

	got := TakeInput("..")
	_, gotKey, gotEntries := model.Json_parser(got)

	if !reflect.DeepEqual(gotKey, wantKey) {
		t.Errorf("got \n%v\nBut want\n%v", gotKey, wantKey)
	}

	if !SameEntries(gotEntries, wantEntries) {
		t.Errorf("got \n%v\nBut want\n%v", gotEntries, wantEntries)
	}
}

func SameEntries(mapA, mapB map[string]model.Entry) bool {
	if len(mapA) != len(mapB) {
		return false
	}
	for key, val := range mapA {
		valB := mapB[key]

		if val != valB {
			return false
		}
	}
	return true
}
