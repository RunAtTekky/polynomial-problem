package helper

import (
	"log"
	"strconv"
)

func BaseConversion(base, value string) int {
	baseInt, err := strconv.Atoi(base)
	if err != nil {
		log.Fatal("Wrong base given", err)
	}

	yValue, err := strconv.ParseInt(value, baseInt, 64)

	return int(yValue)
}
