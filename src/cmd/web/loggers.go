package main

import (
	"io"
	"log"
	"os"
)

// create a logger to log info to stdout and to a file
func (app *application) createInfoLogger() {
	file, err := os.OpenFile("logs/info.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	infoLogger := log.New(multiWriter, "INFO\t", log.Ldate|log.Ltime)
	app.infoLogger = infoLogger
}

// create a logger to log error to stdout and to a file
func (app *application) createErrorLogger() {
	file, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	errorLogger := log.New(multiWriter, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.errorLogger = errorLogger
}
