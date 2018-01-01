package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitWords(t *testing.T) {
	words := splitWords("hello world")
	assert.Equal(t, len(words), 2)
	assert.Equal(t, words[0], "hello")
	assert.Equal(t, words[1], "world")

	words = splitWords("This  is  the  end, Beautiful friend, This is the-end, My only friend, The end, Of our e:laborate plans, the end, Of everything that stands, the end, No safety or surprise, the end. I'll never look into your eyes (again)")
	expectedWords := []string{"this", "is", "the", "end", "beautiful", "friend", "this", "is", "the", "end", "my", "only", "friend", "the", "end", "of", "our", "e", "laborate", "plans", "the", "end", "of", "everything", "that", "stands", "the", "end", "no", "safety", "or", "surprise", "the", "end", "i", "ll", "never", "look", "into", "your", "eyes", "again"}
	assert.Equal(t, words, expectedWords)
}

func Test_stemWords(t *testing.T) {
	words := []string{"the", "runner", "is", "running", "runs"}
	output := stemWords(words)
	assert.Equal(t, output, []string{"the", "runner", "is", "run", "run"})
}

func Test_wordFrequencyCounter(t *testing.T) {
	words := []string{"this", "is", "the", "end", "beautiful", "friend", "this", "is", "the", "end", "my", "only", "friend", "the", "end", "of", "our", "e", "laborate", "plans", "the", "end", "of", "everything", "that", "stands", "the", "end", "no", "safety", "or", "surprise", "the", "end", "i", "ll", "never", "look", "into", "your", "eyes", "again"}
	expectedFrequency := map[string]int{"again": 1, "beautiful": 1, "e": 1, "end": 6, "everything": 1, "eyes": 1, "friend": 2, "i": 1, "into": 1, "is": 2, "laborate": 1, "look": 1, "ll": 1, "my": 1, "never": 1, "no": 1, "of": 2, "only": 1, "or": 1, "our": 1, "plans": 1, "safety": 1, "surprise": 1, "stands": 1, "the": 6, "that": 1, "this": 2, "your": 1}
	frequency := wordFrequencyCounter(words, 1)
	assert.Equal(t, frequency, expectedFrequency)

	words = []string{"one", "one", "two"}
	frequency = wordFrequencyCounter(words, 2)
	assert.Equal(t, frequency, map[string]int{"one": 4, "two": 2})
}

func Test_mergeMaps(t *testing.T) {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	map2 := map[string]int{"two": 2, "three": 3, "four": 4}
	assert.Equal(t, mergeMaps(map1, map2), map[string]int{"one": 1, "two": 4, "three": 6, "four": 4})
}

func Test_buildIndex(t *testing.T) {

	inputSubitems := []subitem{
		subitem{key: "bank", value: "Abhyudaya Cooperative Bank Limited", weight: 5},
		subitem{key: "ifsc", value: "Abhy0065001", weight: 5},
		subitem{key: "micr", value: "400065001", weight: 2},
		subitem{key: "branch", value: "Rtgs-Ho", weight: 2},
		subitem{key: "address", value: "Abhyudaya Bank Bldg., B.No.71, Nehru Nagar, Kurla (E), Mumbai-400024", weight: 1},
		subitem{key: "city", value: "Mumbai", weight: 2},
		subitem{key: "district", value: "Greater Mumbai", weight: 2},
		subitem{key: "state", value: "Maharashtra", weight: 2},
		subitem{key: "contact", value: "25260173", weight: 2},
	}
	input := item{ID: 1, subitems: inputSubitems}
	input.addIndex()

	assert.Equal(t, input.index, map[string]int{"mumbai": 5, "abhyudaya": 6, "nagar": 1, "kurla": 1, "cooper": 5, "rtg": 2, "abhy0065001": 5, "400024": 1, "b": 1, "bldg": 1, "25260173": 2, "greater": 2, "e": 1, "limit": 5, "400065001": 2, "bank": 6, "maharashtra": 2, "71": 1, "no": 1, "nehru": 1, "ho": 2})

}

func Test_globalIndex(t *testing.T) {
	var inx globalIndex
	inx.Init()

	input1 := item{ID: 5, subitems: []subitem{
		subitem{key: "test1", value: "one two three", weight: 1},
	}}
	input1.addIndex()

	input2 := item{ID: 6, subitems: []subitem{
		subitem{key: "test2", value: "two three four", weight: 2},
	}}
	input2.addIndex()

	input3 := item{ID: 7, subitems: []subitem{
		subitem{key: "test3", value: "three four five", weight: 3},
	}}
	input3.addIndex()

	inx.AddLookup(input1)
	inx.AddLookup(input2)
	inx.AddLookup(input3)

	expectedIndex := map[string]map[int]int{
		// On becuase one is stemmed to on
		"on":    map[int]int{5: 1},
		"two":   map[int]int{5: 1, 6: 2},
		"three": map[int]int{5: 1, 6: 2, 7: 3},
		"four":  map[int]int{6: 2, 7: 3},
		"five":  map[int]int{7: 3},
	}
	assert.Equal(t, inx.list, expectedIndex)

}
