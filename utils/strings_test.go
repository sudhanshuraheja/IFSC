package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/ifsc/utils"
)

func Test_SplitWords(t *testing.T) {
	words := utils.SplitWords("hello world")
	assert.Equal(t, len(words), 2)
	assert.Equal(t, words[0], "hello")
	assert.Equal(t, words[1], "world")

	words = utils.SplitWords("This  is  the  end, Beautiful friend, This is the-end, My only friend, The end, Of our e:laborate plans, the end, Of everything that stands, the end, No safety or surprise, the end. I'll never look into your eyes (again)")
	expectedWords := []string{"this", "is", "the", "end", "beautiful", "friend", "this", "is", "the", "end", "my", "only", "friend", "the", "end", "of", "our", "e", "laborate", "plans", "the", "end", "of", "everything", "that", "stands", "the", "end", "no", "safety", "or", "surprise", "the", "end", "i", "ll", "never", "look", "into", "your", "eyes", "again"}
	assert.Equal(t, words, expectedWords)
}

func Test_StemWords(t *testing.T) {
	words := []string{"the", "runner", "is", "running", "runs"}
	output := utils.StemWords(words)
	assert.Equal(t, output, []string{"the", "runner", "is", "run", "run"})
}

func Test_WordFrequencyCounter(t *testing.T) {
	words := []string{"this", "is", "the", "end", "beautiful", "friend", "this", "is", "the", "end", "my", "only", "friend", "the", "end", "of", "our", "e", "laborate", "plans", "the", "end", "of", "everything", "that", "stands", "the", "end", "no", "safety", "or", "surprise", "the", "end", "i", "ll", "never", "look", "into", "your", "eyes", "again"}
	expectedFrequency := map[string]int{"again": 1, "beautiful": 1, "e": 1, "end": 6, "everything": 1, "eyes": 1, "friend": 2, "i": 1, "into": 1, "is": 2, "laborate": 1, "look": 1, "ll": 1, "my": 1, "never": 1, "no": 1, "of": 2, "only": 1, "or": 1, "our": 1, "plans": 1, "safety": 1, "surprise": 1, "stands": 1, "the": 6, "that": 1, "this": 2, "your": 1}
	frequency := utils.WordFrequencyCounter(words, 1)
	assert.Equal(t, frequency, expectedFrequency)

	words = []string{"one", "one", "two"}
	frequency = utils.WordFrequencyCounter(words, 2)
	assert.Equal(t, frequency, map[string]int{"one": 4, "two": 2})
}
