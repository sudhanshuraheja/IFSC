package main

import (
	"fmt"
	"os"

	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/urfave/cli"
)

func main() {
	config.Load()

	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this service lists all bank branches and ifsc codes in india"
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				fmt.Println("Start the service")
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop the service",
			Action: func(c *cli.Context) error {
				fmt.Println("Stop the service")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
