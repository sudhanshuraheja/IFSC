package search

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
