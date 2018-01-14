package excel

import (
	"strings"

	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
	"github.com/sudhanshuraheja/ifsc/model"
)

// AddBanks : fetch banks from the latest excel and upload to db
func AddBanks() error {

	diskFilePath := "./data/latestExcel.xlsx"
	branches := Load(diskFilePath)

	for _, branch := range branches {
		err := addBank(branch)
		if err != nil {
			logger.Fatalln("Could not add bank", branch.Bank, branch.Address)
		}
	}

	logger.Info("Successfully saved all banks to DB")
	return nil
}

func addBank(branch model.Branch) error {
	database := db.Get()

	if branch.Bank == "" {
		logger.Infoln("branch.Bank seems empty, skipping")
		return nil
	}

	if strings.ToLower(branch.Bank) == "bank" {
		logger.Infoln(branch.Bank, "seems invalid, skipping")
		return nil
	}

	// Insert data into DB
	_, err := database.Exec("INSERT INTO branches (bank, ifsc, micr, branch, address, city, district, state, contact) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", branch.Bank, branch.Ifsc, branch.Micr, branch.Branch, branch.Address, branch.City, branch.District, branch.State, branch.Contact)
	if err != nil {
		logger.Fatalln("Pushing to DB failed for", branch.Ifsc, "with error", err)
		return err
	}
	return nil
}
