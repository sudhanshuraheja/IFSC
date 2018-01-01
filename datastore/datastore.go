package datastore

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/sudhanshuraheja/ifsc/search"
)

var inx search.GlobalIndex

// Branch : struct for the data in branch table
type Branch struct {
	DBId      int64  `db:"id"`
	Bank      string `db:"bank"`
	Ifsc      string `db:"ifsc"`
	Micr      string `db:"micr"`
	Branch    string `db:"branch"`
	Address   string `db:"address"`
	City      string `db:"city"`
	District  string `db:"district"`
	State     string `db:"state"`
	Contact   string `db:"contact"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// ToString : convert Branch object to String
func (b Branch) ToString() string {
	return fmt.Sprintf("%s %s %s %s %s %s %s %s %s %s %s %s", strconv.FormatInt(b.DBId, 10), b.Bank, b.Ifsc, b.Micr, b.Branch, b.Address, b.City, b.District, b.State, b.Contact, b.CreatedAt, b.UpdatedAt)
}

// Init : initialise the datastore
func Init() {
	logger.Debugln("Initialising the global search store")
	inx.Init()
}

// ReBuildIndex : build up the search index once again
func ReBuildIndex() {
	logger.Infoln("Got a request to rebuild the index again")

	// Fetch all records from the DB and populate the index
	database := db.Get()
	rows, err := database.Queryx("SELECT * FROM branches LIMIT 2000")
	if err != nil {
		logger.Debugln("Error in query", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b Branch
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

		logger.Infoln("Adding", b.Bank, b.Ifsc, b.Micr, b.Branch, b.Address, b.City, b.District, b.State, b.Contact)

		thisRow.AddIndex()
		inx.AddLookup(thisRow)
	}
}

// Search : search the datastore
func Search(query string) []Branch {
	// return SearchFromPostgres(query)
	return SearchFromGlobalIndex(query)
}

// SearchFromGlobalIndex : search the the new global search index
func SearchFromGlobalIndex(query string) []Branch {
	results := []Branch{}
	database := db.Get()

	escapedQuery, err := url.QueryUnescape(query)
	if err != nil {
		logger.Debugln("Could not parse query", query)
	}
	logger.Debugln("Searching for", escapedQuery)

	ids, err := inx.Find(escapedQuery)
	if err != nil {
		logger.Debugln("Got error from global index", err.Error())
		return results
	}

	foundIds := []string{}
	for ID := range ids {
		logger.Infoln("ID", ID)
		foundIds = append(foundIds, strconv.FormatInt(ID, 10))
	}

	if len(foundIds) == 0 {
		logger.Debug("Found no results in global index")
		return results
	}

	rows, err := database.Queryx("SELECT * FROM branches WHERE id IN (" + strings.Join(foundIds, ",") + ")")

	if err != nil {
		logger.Debugln("Error in query", err)
		return results
	}
	defer rows.Close()

	for rows.Next() {
		var b Branch
		err = rows.StructScan(&b)

		logger.Debugln("DB:", b.ToString())

		if err != nil {
			logger.Debugln("Error is parsing row", err)
		}

		results = append(results, b)
	}

	return results
}

// SearchFromPostgres : search the the postgres db with iLikes
func SearchFromPostgres(query string) []Branch {
	results := []Branch{}
	database := db.Get()

	escapedQuery, err := url.QueryUnescape(query)
	if err != nil {
		logger.Debugln("Could not parse query", query)
	}
	logger.Debugln("Searching for", escapedQuery)

	rows, err := database.Queryx("SELECT * FROM branches WHERE bank ILIKE '%' || $1 || '%' OR ifsc ILIKE '%' || $1 || '%' OR micr ILIKE '%' || $1 || '%' OR branch ILIKE '%' || $1 || '%' OR address ILIKE '%' || $1 || '%' OR city ILIKE '%' || $1 || '%' OR district ILIKE '%' || $1 || '%' OR state ILIKE '%' || $1 || '%' OR contact ILIKE '%' || $1 || '%' LIMIT 10", escapedQuery)
	if err != nil {
		logger.Debugln("Error in query", err)
		return results
	}
	defer rows.Close()

	for rows.Next() {
		var b Branch
		err = rows.StructScan(&b)

		if err != nil {
			logger.Debugln("Error is parsing row", err)
		}

		results = append(results, b)
	}

	return results
}
