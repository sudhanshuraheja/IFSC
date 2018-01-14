package server

import (
	"fmt"
	"net/http"

	"github.com/jeffbmartinez/delay"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"

	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/logger"
)

// StartAPIServer : setup routes and start the server
func StartAPIServer() {
	server := negroni.New()
	router := Router()

	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	if config.EnableDelayMiddleware() {
		server.Use(delay.Middleware{})
	}

	if config.EnableGzipCompression() {
		server.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	if config.EnableStaticFileServer() {
		server.Use(negroni.NewStatic(http.Dir("data")))
	}

	server.Use(Recover())
	server.UseHandler(router)

	serverURL := fmt.Sprintf(":%s", config.Port())
	logger.Infoln("The server is now running at", serverURL)
	http.ListenAndServe(serverURL, server)
}
