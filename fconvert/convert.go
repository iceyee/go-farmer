package fconvert

import (
	"strconv"
	//
)

func S2I(a string) int64 {
	b, _ := strconv.ParseInt(a, 10, 64)
	return b
}

func S2F(a string) float64 {
	b, _ := strconv.ParseFloat(a, 64)
	return b
}

func S2B(a string) bool {
	b, _ := strconv.ParseBool(a)
	return b
}

func I2S(a int64) string {
	return strconv.FormatInt(a, 10)
}

func F2S(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}

func B2S(a bool) string {
	return strconv.FormatBool(a)
}
