package helper

import (
	"log"
	"strconv"
)

func BaseConversion(base, value string) float64 {
	baseInt, err := strconv.Atoi(base)
	if err != nil {
		log.Fatal("Wrong base given", err)
	}

	yValue, err := strconv.ParseInt(value, baseInt, 64)

	return float64(yValue)
}
