package config

import (
	"log"
	"os"

	"github.com/pandulaDW/go-web-project/src/pkg/models/mysql"
)

// Application struct defines the main config for the application
type Application struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	Snippets    *mysql.SnippetModel
}

// CreateInfoLogger creates a logger to log info type logs
func (app *Application) CreateInfoLogger() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	app.InfoLogger = logger
}

// CreateErrorLogger creates a logger to log error type logs
func (app *Application) CreateErrorLogger() {
	logger := log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLogger = logger
}
