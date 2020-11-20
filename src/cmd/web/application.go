package main

import "log"

// using this struct for dependency injection
type application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}
