package excel

import (
	"fmt"
	"strings"
	"time"

	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
)

// UpdateBanks : fetch banks from the latest excel and upload to db
func UpdateBanks() error {
	logger.Infoln("Config", config.LatestDataExcel())

	// diskFilePath := "./data/tmp_ifs_download_1514702975.xlsx"
	diskFilePath := fmt.Sprintf("./data/tmp_ifs_download_%d.xlsx", time.Now().Unix())
	logger.Infoln("Save to disk at", diskFilePath)

	err := download(config.LatestDataExcel(), diskFilePath)
	if err != nil {
		logger.Fatal(err)
	}

	branches := Load(diskFilePath)

	database := db.Get()
	for _, branch := range branches {

		if branch.Bank == "" {
			logger.Infoln("branch.Bank seems empty, skipping")
			continue
		}

		if strings.ToLower(branch.Bank) == "bank" {
			logger.Infoln(branch.Bank, "seems invalid, skipping")
			continue
		}

		// Insert data into DB
		_, err := database.Exec("INSERT INTO branches (bank, ifsc, micr, branch, address, city, district, state, contact) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", branch.Bank, branch.Ifsc, branch.Micr, branch.Branch, branch.Address, branch.City, branch.District, branch.State, branch.Contact)
		if err != nil {
			logger.Fatalln("Pushing to DB failed for", branch.Ifsc, "with error", err)
			return err
		}
	}

	logger.Info("Successfull saved everything to DB")
	return nil
}
