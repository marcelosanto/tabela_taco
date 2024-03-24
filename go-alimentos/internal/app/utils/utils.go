package utils

import (
	"math"
	"strconv"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 1, 64)
}

func FloatRoundToString(num float64) string {
	numRound := int(math.Round(num))
	return strconv.Itoa(numRound)
}

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}
