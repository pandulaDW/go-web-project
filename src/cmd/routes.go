package main

import (
	"net/http"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
	"github.com/pandulaDW/go-web-project/src/cmd/snippets"
)

func routes(app *config.Application) *http.ServeMux {
	// create the serve mux router
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("/", snippets.Home(app))
	mux.HandleFunc("/snippet", snippets.ShowSnippet(app))
	mux.HandleFunc("/snippets", snippets.LatestSnippets(app))
	mux.HandleFunc("/snippet/create", snippets.CreateSnippet(app))

	// static file handling
	fileServer := http.FileServer(http.Dir("./src/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
