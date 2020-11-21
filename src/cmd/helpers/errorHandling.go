package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
)

// ServeError will print a stack trace and return the error to the user
func ServeError(w http.ResponseWriter, err error, app *config.Application) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// ClientError will return the appropriate client error to the user
func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound will return a NotFound error to the user
func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
