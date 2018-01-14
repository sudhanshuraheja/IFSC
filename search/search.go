package search

import (
	"github.com/sudhanshuraheja/ifsc/utils"
)

// Item : each item that needs to be indexed
type Item struct {
	ID       int64
	SubItems []SubItem
	Index    map[string]int
}

// SubItem : each component inside an item that can have a separate weight
type SubItem struct {
	Key    string
	Value  string
	Weight int
}

// AddIndex : add index for a specific item
func (i *Item) AddIndex() {
	i.Index = make(map[string]int)
	for _, x := range i.SubItems {
		i.Index = utils.MergeMaps(i.Index, addSubItemIndex(x.Value, x.Weight))
	}
}

// AddSubItem : Add a subItem from an external file
func addSubItemIndex(value string, weight int) map[string]int {
	words := utils.SplitWords(value)
	frequency := wordFrequencyCounter(utils.StemWords(words), weight)
	return frequency
}
