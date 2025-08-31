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
	wantRawData, _, _ := model.Json_parser(want)

	got := TakeInput("..")
	gotRawData, _, _ := model.Json_parser(got)

	if !reflect.DeepEqual(wantRawData, gotRawData) {
		t.Errorf("got \n%v\nBut want\n%v", got, want)
	}
}
