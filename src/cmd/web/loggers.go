package main

import (
	"log"
	"os"
)

func (app *application) createInfoLogger() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	app.infoLogger = logger
}

func (app *application) createErrorLogger() {
	logger := log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	app.errorLogger = logger
}
