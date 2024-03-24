package utils

import (
	"math"
	"strconv"
)

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func FloatRoundToString(num float64) string {
	numRound := int(math.Round(num))

	return strconv.Itoa(numRound)
}
