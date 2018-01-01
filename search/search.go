package search

import (
	"strings"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

type item struct {
	ID       int
	subitems []subitem
	index    map[string]int
}

type subitem struct {
	key    string
	value  string
	weight int
}

func (i *item) addIndex() {
	i.index = make(map[string]int)
	for _, x := range i.subitems {
		i.index = mergeMaps(i.index, addSubItemIndex(x.value, x.weight))
	}
}

// AddSubItem : Add a subItem from an external file
func addSubItemIndex(value string, weight int) map[string]int {
	words := splitWords(value)
	frequency := wordFrequencyCounter(stemWords(words), weight)
	return frequency
}

func splitWords(line string) []string {
	line = strings.ToLower(line)
	splitList := []string{" ", ",", ":", "-", "(", ")", ".", "'"}
	words := strings.FieldsFunc(line, func(r rune) bool {
		success := false
		for _, split := range splitList {
			if string(r) == split {
				success = true
			}
		}
		return success
	})
	return words
}

func stemWords(words []string) []string {
	output := []string{}
	for _, word := range words {
		output = append(output, porterstemmer.StemString(word))
	}
	return output
}

func wordFrequencyCounter(words []string, weight int) map[string]int {
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

func mergeMaps(map1 map[string]int, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	for key, value := range map1 {
		result[key] = value
	}

	for key := range map2 {
		_, keyExists := map2[key]
		if keyExists {
			result[key] = result[key] + map2[key]
		} else {
			result[key] = map2[key]
		}
	}

	return result
}
