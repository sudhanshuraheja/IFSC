package cli

import (
	"os"

	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/db"
	"github.com/sudhanshuraheja/ifsc/excel"
	"github.com/sudhanshuraheja/ifsc/server"
	"github.com/urfave/cli"
)

// Start : start the cli wrapper
func Start() {
	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this service lists all bank branches and ifsc codes in india"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return db.RunDatabaseMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return db.RollbackDatabaseMigration()
			},
		},
		{
			Name:  "updateBanks",
			Usage: "take the latest list of banks and update db",
			Action: func(c *cli.Context) error {
				return excel.UpdateBanks()
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
