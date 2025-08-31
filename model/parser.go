package model

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
