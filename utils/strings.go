package utils

import (
	"strings"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

var separators = []string{" ", ",", ":", "-", "(", ")", ".", "'"}

// SplitWords takes a string and breaks it up into words, based on separators
func SplitWords(line string) []string {
	line = strings.ToLower(line)
	words := strings.FieldsFunc(line, func(r rune) bool {
		success := false
		for _, split := range separators {
			if string(r) == split {
				success = true
			}
		}
		return success
	})
	return words
}

// StemWord returns the stemmed word
func StemWord(query string) string {
	return porterstemmer.StemString(query)
}

// StemWords with porterStemmer
func StemWords(words []string) []string {
	output := []string{}
	for _, word := range words {
		output = append(output, porterstemmer.StemString(word))
	}
	return output
}

// WordFrequencyCounter counts the frequency of each word
func WordFrequencyCounter(words []string, weight int) map[string]int {
	frequencyCounter := make(map[string]int)

	for _, word := range words {
		_, wordExists := frequencyCounter[word]
		if wordExists {
			frequencyCounter[word] = frequencyCounter[word] + (1 * weight)
		} else {
			frequencyCounter[word] = (1 * weight)
		}
	}

	return frequencyCounter
}
