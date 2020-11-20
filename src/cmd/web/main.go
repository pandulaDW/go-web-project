package main

import (
	"flag"
	"net/http"
)

// define global app variable
var app application

func init() {
	// create a new app instance
	app = application{}
}

func main() {
	// create the loggers
	app.createInfoLogger()
	app.createErrorLogger()

	// reading configuration info from command line
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// static file handling
	fileServer := http.FileServer(http.Dir("./src/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// create a new http server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLogger.logger,
		Handler:  mux,
	}

	app.infoLogger.PrintLog("Starting server on " + *addr)
	err := srv.ListenAndServe()
	app.errorLogger.logger.Fatal(err)
}
