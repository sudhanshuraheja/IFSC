package datastore

import (
	"fmt"
	"strings"

	"github.com/sudhanshuraheja/ifsc/excel"
	"github.com/sudhanshuraheja/ifsc/logger"
)

type store struct {
	branches excel.Branches
}

var branchStore store

// Init : initialise the datastore
func Init(dataFile string) {
	branchStore.branches = excel.Load(dataFile)
	logger.Infoln("Loaded up data for", branchStore.branches.Count, "banks")
}

// Search : search the datastore
func Search(query string) []excel.Branch {
	results := []excel.Branch{}
	maxResults := 10

	if query == "" {
		return results
	}

	for _, bank := range branchStore.branches.List {
		if len(results) >= maxResults {
			break
		}
		mergedSting := fmt.Sprintf("%s %s %s %s %s %s %s %s %s", bank.Branch, bank.Ifsc, bank.Micr, bank.Branch, bank.Address, bank.Contact, bank.City, bank.District, bank.State)
		if strings.Contains(strings.ToLower(mergedSting), strings.ToLower(query)) {
			results = append(results, bank)
		}
	}

	return results
}
