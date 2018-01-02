package search

import (
	"errors"
	"fmt"

	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/sudhanshuraheja/ifsc/model"

	porterstemmer "github.com/reiver/go-porterstemmer"
)

// WordIndex : struct for indexes in the database
type WordIndex struct {
	ID       int64  `db:"id"`
	Key      string `db:"key"`
	BranchID int64  `db:"branch"`
	Weight   int    `db:"weight"`
}

func (w *WordIndex) toString() string {
	return fmt.Sprintln("[", w.ID, "]", w.Key, w.BranchID, w.Weight)
}

// AddLookup : add indexes for a particular item to the index
func (w *WordIndex) AddLookup(i Item) error {
	branchid := i.ID
	database := db.Get()

	for key, weight := range i.Index {

		_, err := database.Exec("INSERT INTO search (key, branch, weight) VALUES ($1, $2, $3) ON CONFLICT (key, branch) DO UPDATE SET weight = $3 WHERE search.key = $1 and search.branch = $2", key, branchid, weight)
		if err != nil {
			logger.Fatalln("Pushing to DB for", key, branchid, weight, "failed with error", err)
			return err
		}
	}
	return nil
}

func (w *WordIndex) findKey(query string) (string, error) {
	stemmedKey := porterstemmer.StemString(query)
	return stemmedKey, nil
}

// Find : find list of items which match this query
func (w *WordIndex) Find(query string) ([]model.Branch, error) {
	results := []model.Branch{}
	key, err := w.findKey(query)
	if err != nil {
		return []model.Branch{}, errors.New("We could not find any search results for " + query)
	}

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

		logger.Debugln("ROW:", b.ToString())

		if err != nil {
			logger.Debugln("Error is parsing row", err)
		}

		results = append(results, b)
	}

	// return g.list[key], nil
	return results, nil
}
