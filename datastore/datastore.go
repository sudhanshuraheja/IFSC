package datastore

import (
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
	maxResults := 5

	if query == "" {
		return results
	}

	for _, bank := range branchStore.branches.List {
		if len(results) >= maxResults {
			break
		}
		if strings.Contains(strings.ToLower(bank.FullText), strings.ToLower(query)) {
			results = append(results, bank)
		}
	}

	return results
}
