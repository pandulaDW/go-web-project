package main

import (
	"flag"
	"net/http"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
)

func main() {
	// define initial app
	app := config.Application{}

	// create the loggers
	app.CreateInfoLogger()
	app.CreateErrorLogger()

	// reading configuration info from command line
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// create a new http server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLogger,
		Handler:  routes(&app),
	}

	// starting the server
	app.InfoLogger.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	app.ErrorLogger.Fatal(err)
}
