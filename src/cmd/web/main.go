package main

import (
	"flag"
	"net/http"
)

func main() {
	// create a new app instance
	app := application{}

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
		ErrorLog: app.errorLogger,
		Handler:  mux,
	}

	app.infoLogger.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	app.errorLogger.Fatal(err)
}
