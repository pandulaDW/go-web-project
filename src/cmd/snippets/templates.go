package snippets

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/pandulaDW/go-web-project/src/pkg/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func createTemplateWithDefaults() *templateData {
	td := &templateData{}
	td.CurrentYear = time.Now().Year()
	return td
}

// NewTemplateCache creates a cache for templates
func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.htm"))
	if err != nil {
		return nil, err
	}

	// Loop through the pages one-by-one.
	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.htm"))
		if err != nil {
			return nil, err
		}

		// ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.htm"))
		// if err != nil {
		// 	return nil, err
		// }

		cache[name] = ts
	}

	return cache, nil
}
