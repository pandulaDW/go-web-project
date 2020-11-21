package snippets

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/pandulaDW/go-web-project/src/cmd/config"
)

// Home returns the handler function for Home route
func Home(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./src/ui/html/home.page.htm",
			"./src/ui/html/footer.layout.htm",
			"./src/ui/html/base.layout.htm",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLogger.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLogger.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}

	return handler
}

// ShowSnippet returns the handler function for ShowSnippet route
func ShowSnippet(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "Alex", "id": %d}`, id)
	}

	return handler
}

// CreateSnippet returns the handler function for ShowSnippet route
func CreateSnippet(app *config.Application) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Method not allowed", 405)
			return
		}
		w.Write([]byte("Create a new snippet..."))
	}

	return handler
}