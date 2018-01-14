package search

import (
	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/sudhanshuraheja/ifsc/model"
	"github.com/sudhanshuraheja/ifsc/utils"
)

// BuildIndex : build up the search index
func BuildIndex() {
	logger.Infoln("Got a request to rebuild the index again")
	getBranches()
}

func getBranches() {
	db := db.Get()
	// Fetch all records from the DB and populate the index
	rows, err := db.Queryx("SELECT * FROM branches")
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

		keywords := make(map[string]int)
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.Bank, 5))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.Ifsc, 3))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.Micr, 1))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.Branch, 3))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.Address, 1))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.City, 2))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.District, 2))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.State, 2))
		keywords = utils.MergeMaps(keywords, addBankIndexKeywords(b.Contact, 2))

		saveKeywords(b.DBId, keywords)
	}
}

func addBankIndexKeywords(value string, weight int) map[string]int {
	words := utils.SplitWords(value)
	frequency := utils.WordFrequencyCounter(utils.StemWords(words), weight)
	return frequency
}

func saveKeywords(branch int64, keywords map[string]int) error {
	database := db.Get()

	for key, weight := range keywords {

		_, err := database.Exec("INSERT INTO search (branch, key, weight) VALUES ($1, $2, $3) ON CONFLICT (branch, key) DO UPDATE SET weight = $3 WHERE search.branch = $1 and search.key = $2", branch, key, weight)
		if err != nil {
			logger.Fatalln("Pushing to DB for", branch, key, weight, "failed with error", err)
			return err
		}
	}
	return nil
}
