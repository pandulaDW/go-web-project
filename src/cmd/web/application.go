package main

// using this struct for dependency injection
type application struct {
	errorLogger *multiLogger
	infoLogger  *multiLogger
}
