package helpers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
)

// Render will render the html templates. It will first write into a bytes buffer and if it
// fails, it will return an error to the receiver
func Render(w http.ResponseWriter, r *http.Request, name string, td interface{}, app *config.Application) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		ServeError(w, fmt.Errorf("The template %s does not exist", name), app)
		return
	}

	// initialize a new buffer
	buf := new(bytes.Buffer)

	err := ts.Execute(buf, td)
	if err != nil {
		ServeError(w, err, app)
		return
	}

	// write the contents of the buffer to the http.ResponseWriter
	buf.WriteTo(w)
}
