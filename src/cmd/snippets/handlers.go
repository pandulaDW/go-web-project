package snippets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
	"github.com/pandulaDW/go-web-project/src/cmd/helpers"
	"github.com/pandulaDW/go-web-project/src/pkg/models"
)

// Home returns the handler function for Home route
func Home(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			helpers.NotFound(w)
			return
		}

		results, err := app.Snippets.Latest()
		if err != nil {
			helpers.ServeError(w, err, app)
		}

		dataModel := createTemplateWithDefaults()
		dataModel.Snippets = results
		helpers.Render(w, r, "home.page.htm", dataModel, app)
	}

	return handler
}

// ShowSnippet returns the handler function for ShowSnippet route
func ShowSnippet(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// validate id is a positive integer
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			helpers.NotFound(w)
			return
		}

		s, err := app.Snippets.Get(id)
		if err == models.ErrNoRecord {
			helpers.NotFound(w)
			return
		} else if err != nil {
			helpers.ServeError(w, err, app)
			return
		}

		dataModel := createTemplateWithDefaults()
		dataModel.Snippet = s
		helpers.Render(w, r, "show.page.htm", dataModel, app)
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
