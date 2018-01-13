package datastore

import (
	"net/url"

	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/sudhanshuraheja/ifsc/model"
	"github.com/sudhanshuraheja/ifsc/search"
)

var inx search.WordIndex

// ReBuildIndex : build up the search index once again
func ReBuildIndex() {
	logger.Infoln("Got a request to rebuild the index again")

	// Fetch all records from the DB and populate the index
	database := db.Get()
	rows, err := database.Queryx("SELECT * FROM branches")
	if err != nil {
		logger.Debugln("Error in query", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b model.Branch
		err = rows.StructScan(&b)

		if err != nil {
			logger.Debugln("Error is parsing row", err)
		}

		thisRow := search.Item{ID: b.DBId, SubItems: []search.SubItem{
			search.SubItem{Key: "bank", Value: b.Bank, Weight: 5},
			search.SubItem{Key: "ifsc", Value: b.Ifsc, Weight: 3},
			search.SubItem{Key: "micr", Value: b.Micr, Weight: 1},
			search.SubItem{Key: "branch", Value: b.Branch, Weight: 3},
			search.SubItem{Key: "address", Value: b.Address, Weight: 1},
			search.SubItem{Key: "city", Value: b.City, Weight: 2},
			search.SubItem{Key: "district", Value: b.District, Weight: 2},
			search.SubItem{Key: "state", Value: b.State, Weight: 2},
			search.SubItem{Key: "contact", Value: b.Contact, Weight: 2},
		}}

		thisRow.AddIndex()
		inx.AddLookup(thisRow)
	}
}

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

	results, err := inx.Find(escapedQuery)
	if err != nil {
		logger.Debugln("Caught error", err.Error())
		return []model.Branch{}
	}

	return results
}
