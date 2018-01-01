package search

import (
	"errors"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

type globalIndex struct {
	// map[key]lookup
	// 		   map[ID]weight
	list map[string]map[int]int
}

func (g *globalIndex) Init() {
	g.list = make(map[string]map[int]int)
}

func (g *globalIndex) AddLookup(i item) {
	id := i.ID
	for key, weight := range i.index {
		// Check if key exists in globalIndex
		_, keyExists := g.list[key]
		if keyExists {
			g.list[key][id] = weight
		} else {
			// Key doesn't exist, add key, id, weight map
			weightForID := make(map[int]int)
			weightForID[id] = weight
			g.list[key] = weightForID
		}
	}
}

func (g *globalIndex) findKey(query string) (string, error) {
	stemmedKey := porterstemmer.StemString(query)
	_, keyExists := g.list[stemmedKey]
	if !keyExists {
		return "", errors.New("Could not find key " + query)
	}
	return stemmedKey, nil
}

func (g *globalIndex) Find(query string) (map[int]int, error) {
	key, err := g.findKey(query)
	if err != nil {
		return map[int]int{}, errors.New("We could not find any search results for " + query)
	}

	return g.list[key], nil
}
