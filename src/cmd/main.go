package main

import (
	"flag"
	"net/http"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
	"github.com/pandulaDW/go-web-project/src/cmd/snippets"
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

	// create the serve mux router
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("/", snippets.Home(&app))
	mux.HandleFunc("/snippet", snippets.ShowSnippet(&app))
	mux.HandleFunc("/snippet/create", snippets.CreateSnippet(&app))

	// static file handling
	fileServer := http.FileServer(http.Dir("./src/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// create a new http server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLogger,
		Handler:  mux,
	}

	// starting the server
	app.InfoLogger.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	app.ErrorLogger.Fatal(err)
}
