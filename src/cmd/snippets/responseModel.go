package snippets

import (
	"fmt"

	"github.com/pandulaDW/go-web-project/src/pkg/models"
)

func showSnippetResponse(s *models.Snippet) string {
	response := fmt.Sprintf(`
	{ "id": %d,"Title": "%s","Content": "%s","Created": %v,"Expires": %v}`,
		s.ID, s.Title, s.Content, s.Created, s.Expires)
	return response
}
