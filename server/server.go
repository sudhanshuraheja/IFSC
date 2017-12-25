package server

import (
	"net/http"

	"github.com/jeffbmartinez/delay"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

// StartAPIServer : setup routes and start the server
func StartAPIServer() {
	server := negroni.New()
	router := Router()

	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())
	server.Use(delay.Middleware{})
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	// TODO : Move this behind remote config
	// Usually not required
	// server.Use(negroni.NewStatic(http.Dir("public")))
	server.Use(Recover())
	server.UseHandler(router)

	http.ListenAndServe(":3000", server)
}
