package main

import (
	"github.com/sudhanshuraheja/ifsc/cli"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/datastore"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func main() {
	config.Load()
	logger.Setup()
	datastore.Init("data/sample.xlsx")
	// db.Init()
	cli.Start()
}
