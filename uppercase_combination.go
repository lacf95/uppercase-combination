package uppercase_combination

import (
	"math"
	"strings"
	"unicode"
)

func UppercaseCombination(str string) []string {
	length := len(str)
	combinations := kCombination(length)
	result := stringSlice(combinations, str)

	for i := 0; i < length; i++ {
		interval := combinationInterval(combinations, i + 1)
		counter := interval
		for counter < combinations {
			result[counter] = uppercaseRuneAt(result[counter], i)

			counter++
			if counter%interval == 0 {
				counter = counter + interval
			}
		}
	}
	return result
}

func kCombination(n int) int {
	return int(math.Pow(float64(2), float64(n)))
}

func combinationInterval(n, i int) int {
	return int(float64(n) / math.Pow(float64(2), float64(i)))
}

func stringSlice(length int, str string) []string {
	result := make([]string, length)
	baseForm := strings.ToLower(str)
	for i := 0; i < length; i++ {
		result[i] = baseForm
	}
	return result
}

func uppercaseRuneAt(str string, index int) string {
	runes := []rune(str)
	runes[index] = unicode.ToUpper(runes[index])
	return string(runes)
}
