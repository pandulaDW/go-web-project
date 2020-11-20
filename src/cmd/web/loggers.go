package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// create a logger to log info to stdout and to a file
func (app *application) createInfoLogger() {
	now := time.Now()
	today := fmt.Sprintf("%d_%d_%d", now.Year(), now.Month(), now.Day())
	filePath := fmt.Sprintf("logs/info_%s.txt", today)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	infoLogger := log.New(multiWriter, "INFO\t", log.Ldate|log.Ltime)
	app.infoLogger = infoLogger
}

// create a logger to log error to stdout and to a file
func (app *application) createErrorLogger() {
	now := time.Now()
	today := fmt.Sprintf("%d_%d_%d", now.Year(), now.Month(), now.Day())
	filePath := fmt.Sprintf("logs/error_%s.txt", today)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}

	multiWriter := io.MultiWriter(file, os.Stderr)

	errorLogger := log.New(multiWriter, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.errorLogger = errorLogger
}
