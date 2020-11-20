package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type logType string

const (
	// INFO type
	INFO logType = "INFO"
	// ERROR type
	ERROR logType = "ERROR"
)

type multiLogger struct {
	now     time.Time
	today   string
	logType logType
	logger  *log.Logger
}

func createDateFormat(givenTime time.Time) string {
	return fmt.Sprintf("%d_%d_%d", givenTime.Year(), givenTime.Month(), givenTime.Day())
}

func (customLogger *multiLogger) assignDate() {
	now := time.Now()
	customLogger.now = now
	customLogger.today = createDateFormat(now)
}

// PrintLog prints message according to the relevant logger. If it's a new day,
// it will create a new logger that writes to a new log file
func (customLogger *multiLogger) PrintLog(message string) {
	if createDateFormat(customLogger.now) != createDateFormat(time.Now()) {
		customLogger.assignDate()

		switch customLogger.logType {
		case INFO:
			app.createInfoLogger()
		case ERROR:
			app.createErrorLogger()
		}
	}
	customLogger.logger.Println(message)
}

// create a logger to log info to stdout and to a file
func (app *application) createInfoLogger() {
	if app.infoLogger == nil {
		app.infoLogger = &multiLogger{}
		app.infoLogger.logType = INFO
		app.infoLogger.assignDate()
	}
	filePath := fmt.Sprintf("logs/info_%s.txt", app.infoLogger.today)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	multiWriter := io.MultiWriter(file, os.Stdout)
	app.infoLogger.logger = log.New(multiWriter, "INFO\t", log.Ldate|log.Ltime)
}

// create a logger to log error to stderr and to a file
func (app *application) createErrorLogger() {
	if app.errorLogger == nil {
		app.errorLogger = &multiLogger{}
		app.errorLogger.logType = ERROR
		app.errorLogger.assignDate()
	}
	filePath := fmt.Sprintf("logs/error_%s.txt", app.errorLogger.today)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	multiWriter := io.MultiWriter(file, os.Stderr)
	app.errorLogger.logger = log.New(multiWriter, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
