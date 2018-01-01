package main

import (
	"github.com/sudhanshuraheja/ifsc/cli"
	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/datastore"
	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/logger"
)

func main() {
	config.Load()
	logger.Setup()
	db.Init()
	defer db.Close()
	cli.Start()
	datastore.Init()
}
