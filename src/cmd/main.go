package main

import (
	"database/sql"
	"flag"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pandulaDW/go-web-project/src/cmd/config"
	"github.com/pandulaDW/go-web-project/src/pkg/models/mysql"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

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

	// connect to db
	db, err := openDB("pandula:gotohell2@/snippetbox?parseTime=true")
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}
	defer db.Close()

	// initialize snippet model
	app.Snippets = &mysql.SnippetModel{DB: db}

	// starting the server
	app.InfoLogger.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	app.ErrorLogger.Fatal(err)
}
