package datastore

import (
	"net/url"

	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/sudhanshuraheja/ifsc/model"
	"github.com/sudhanshuraheja/ifsc/search"
)

// Search : search the datastore
func Search(query string) []model.Branch {
	// return SearchFromPostgres(query)
	return SearchFromIndex(query)
}

// SearchFromIndex : search the the new db search index
func SearchFromIndex(query string) []model.Branch {
	escapedQuery, err := url.QueryUnescape(query)
	if err != nil {
		logger.Debugln("Could not parse query", query)
		return []model.Branch{}
	}
	logger.Debugln("Searching for", escapedQuery)

	results, err := search.Find(escapedQuery)
	if err != nil {
		logger.Debugln("Caught error", err.Error())
		return []model.Branch{}
	}

	return results
}
