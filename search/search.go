package search

import (
	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/sudhanshuraheja/ifsc/model"
	"github.com/sudhanshuraheja/ifsc/utils"
)

// Find : find list of items which match this query
func Find(query string) ([]model.Branch, error) {
	results := []model.Branch{}
	key := utils.StemWord(query)

	database := db.Get()
	rows, err := database.Queryx("SELECT * FROM branches WHERE id IN (SELECT branch FROM SEARCH WHERE key=$1 ORDER BY weight DESC LIMIT 25)", key)

	if err != nil {
		logger.Debugln("Error in query", err)
		return results, nil
	}
	defer rows.Close()

	for rows.Next() {
		var b model.Branch
		err = rows.StructScan(&b)

		if err != nil {
			logger.Debugln("Error is parsing row", err)
		}

		results = append(results, b)
	}

	return results, nil
}
