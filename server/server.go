package server

import (
	"github.com/urfave/negroni"
)

// StartAPIServer : setup routes and start the server
func StartAPIServer() {
	server := negroni.New(negroni.NewRecovery())
	router := Router()

	server.Use(Recover())
	server.UseHandler(router)
	server.Run(":3000")
}
