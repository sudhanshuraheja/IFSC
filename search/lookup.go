package search

import (
	"errors"
	"fmt"
	"strconv"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

// GlobalIndex : list of all indexes
type GlobalIndex struct {
	// map[key]lookup
	// 		   map[ID]weight
	list map[string]map[int64]int
}

// Init : initialise the global index
func (g *GlobalIndex) Init() {
	g.list = make(map[string]map[int64]int)
}

func (g *GlobalIndex) info() {
	linkages := 0
	for _, value := range g.list {
		linkages += len(value)
	}
	fmt.Println("Index now has " + strconv.Itoa(len(g.list)) + " keys and " + strconv.Itoa(linkages) + " linkages")
}

// AddLookup : add indexes for a particular item to the index
func (g *GlobalIndex) AddLookup(i Item) {
	id := i.ID
	for key, weight := range i.Index {
		// Check if key exists in globalIndex
		_, keyExists := g.list[key]
		if keyExists {
			g.list[key][id] = weight
		} else {
			// Key doesn't exist, add key, id, weight map
			weightForID := make(map[int64]int)
			weightForID[id] = weight
			g.list[key] = weightForID
		}
	}
	g.info()
}

func (g *GlobalIndex) findKey(query string) (string, error) {
	stemmedKey := porterstemmer.StemString(query)
	_, keyExists := g.list[stemmedKey]
	if !keyExists {
		return "", errors.New("Could not find key " + query)
	}
	return stemmedKey, nil
}

// Find : find list of items which match this query
func (g *GlobalIndex) Find(query string) (map[int64]int, error) {
	key, err := g.findKey(query)
	if err != nil {
		return map[int64]int{}, errors.New("We could not find any search results for " + query)
	}

	return g.list[key], nil
}
