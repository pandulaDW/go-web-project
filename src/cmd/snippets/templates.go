package snippets

import "github.com/pandulaDW/go-web-project/src/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
