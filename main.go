package main

import (
	"github.com/sudhanshuraheja/ifsc/cli"
	"github.com/sudhanshuraheja/ifsc/config"
)

func main() {
	config.Load()
	cli.Start()
}
