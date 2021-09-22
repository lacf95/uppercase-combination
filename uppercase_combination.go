package uppercase_combination

import (
	"math"
	"strings"
	"unicode"
)

func UppercaseCombination(str string) []string {
	length := len(str)
	baseForm := strings.ToLower(str)
	combinationsLength := kCombination(length)
	combinations := make([]string, combinationsLength)
	combinations[0] = baseForm

	for i := 0; i < length; i++ {
		interval := combinationInterval(combinationsLength, i + 1)
		counter := interval
		for counter < combinationsLength {
			if combinations[counter] == "" {
				combinations[counter] = baseForm
			}
			combinations[counter] = uppercaseRuneAt(combinations[counter], i)

			counter++
			if counter%interval == 0 {
				counter = counter + interval
			}
		}
	}
	return combinations
}

func kCombination(n int) int {
	return int(math.Pow(float64(2), float64(n)))
}

func combinationInterval(n, i int) int {
	return int(float64(n) / math.Pow(float64(2), float64(i)))
}

func uppercaseRuneAt(str string, index int) string {
	runes := []rune(str)
	runes[index] = unicode.ToUpper(runes[index])
	return string(runes)
}
