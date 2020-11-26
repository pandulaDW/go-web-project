package helpers

import (
	"fmt"
	"net/http"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
)

func Render(w http.ResponseWriter, r *http.Request, name string, td interface{}, app *config.Application) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		ServeError(w, fmt.Errorf("The template %s does not exist", name), app)
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
		ServeError(w, err, app)
	}
}
