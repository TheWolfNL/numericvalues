package numericvalues

import "strings"

func combineArrays(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}
	return r
}

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var numericalValues = []string{"10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35"}

func LettersToNumericValue(str string) string {
	return strings.NewReplacer(combineArrays(alphabet, numericalValues)...).Replace(str)
}

func NumericValueToLetters(str string) string {
	return strings.NewReplacer(combineArrays(numericalValues, alphabet)...).Replace(str)
}
