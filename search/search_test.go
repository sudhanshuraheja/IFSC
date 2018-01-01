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

	inputSubitems := []SubItem{
		SubItem{Key: "bank", Value: "Abhyudaya Cooperative Bank Limited", Weight: 5},
		SubItem{Key: "ifsc", Value: "Abhy0065001", Weight: 5},
		SubItem{Key: "micr", Value: "400065001", Weight: 2},
		SubItem{Key: "branch", Value: "Rtgs-Ho", Weight: 2},
		SubItem{Key: "address", Value: "Abhyudaya Bank Bldg., B.No.71, Nehru Nagar, Kurla (E), Mumbai-400024", Weight: 1},
		SubItem{Key: "city", Value: "Mumbai", Weight: 2},
		SubItem{Key: "district", Value: "Greater Mumbai", Weight: 2},
		SubItem{Key: "state", Value: "Maharashtra", Weight: 2},
		SubItem{Key: "contact", Value: "25260173", Weight: 2},
	}
	input := Item{ID: 1, SubItems: inputSubitems}
	input.AddIndex()

	assert.Equal(t, input.Index, map[string]int{"mumbai": 5, "abhyudaya": 6, "nagar": 1, "kurla": 1, "cooper": 5, "rtg": 2, "abhy0065001": 5, "400024": 1, "b": 1, "bldg": 1, "25260173": 2, "greater": 2, "e": 1, "limit": 5, "400065001": 2, "bank": 6, "maharashtra": 2, "71": 1, "no": 1, "nehru": 1, "ho": 2})

}

func Test_globalIndex(t *testing.T) {
	var inx GlobalIndex
	inx.Init()

	input1 := Item{ID: 5, SubItems: []SubItem{
		SubItem{Key: "test1", Value: "one two three", Weight: 1},
	}}
	input1.AddIndex()

	input2 := Item{ID: 6, SubItems: []SubItem{
		SubItem{Key: "test2", Value: "two three four", Weight: 2},
	}}
	input2.AddIndex()

	input3 := Item{ID: 7, SubItems: []SubItem{
		SubItem{Key: "test3", Value: "three four five", Weight: 3},
	}}
	input3.AddIndex()

	inx.AddLookup(input1)
	inx.AddLookup(input2)
	inx.AddLookup(input3)

	expectedIndex := map[string]map[int64]int{
		// On becuase one is stemmed to on
		"on":    map[int64]int{5: 1},
		"two":   map[int64]int{5: 1, 6: 2},
		"three": map[int64]int{5: 1, 6: 2, 7: 3},
		"four":  map[int64]int{6: 2, 7: 3},
		"five":  map[int64]int{7: 3},
	}
	assert.Equal(t, inx.list, expectedIndex)

	ids, err := inx.Find("three")
	assert.Equal(t, err, nil)
	assert.Equal(t, ids, map[int64]int{5: 1, 6: 2, 7: 3})

	ids, err = inx.Find("six")
	assert.Equal(t, err.Error(), "We could not find any search results for six")
	assert.Equal(t, ids, map[int64]int{})

}
