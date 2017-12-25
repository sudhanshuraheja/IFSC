package datastore

import (
	"github.com/sudhanshuraheja/ifsc/excel"
	"github.com/sudhanshuraheja/ifsc/logger"
)

// Init : initialise the datastore
func Init() {
	allBranches := excel.Load("data/sample.xlsx")
	logger.Infoln("Loaded up data for", allBranches.Count, "banks")
}
