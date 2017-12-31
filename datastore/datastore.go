package datastore

import (
	"net/url"

	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
)

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

// Search : search the datastore
func Search(query string) []Branch {
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
