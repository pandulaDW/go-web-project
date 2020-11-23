package snippets

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
	"github.com/pandulaDW/go-web-project/src/cmd/helpers"
)

// Home returns the handler function for Home route
func Home(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			helpers.NotFound(w)
			return
		}

		files := []string{
			"./src/ui/html/home.page.htm",
			"./src/ui/html/footer.layout.htm",
			"./src/ui/html/base.layout.htm",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			helpers.ServeError(w, err, app)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			helpers.ServeError(w, err, app)
		}
	}

	return handler
}

// ShowSnippet returns the handler function for ShowSnippet route
func ShowSnippet(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			helpers.NotFound(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "New Snippet", "id": %d}`, id)
	}

	return handler
}

// CreateSnippet returns the handler function for ShowSnippet route
func CreateSnippet(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			helpers.ClientError(w, http.StatusMethodNotAllowed)
			return
		}

		// dummy data
		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
		expires := "7"

		id, err := app.Snippets.Insert(title, content, expires)
		if err != nil {
			helpers.ServeError(w, err, app)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	}

	return handler
}
