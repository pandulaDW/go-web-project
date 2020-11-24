package snippets

import (
	"encoding/json"

	"github.com/pandulaDW/go-web-project/src/pkg/models"
)

func showSnippetResponse(s *models.Snippet) []byte {
	response, _ := json.Marshal(s)
	return response
}

func latestSnippetResponse(s []*models.Snippet) []byte {
	response, _ := json.Marshal(s)
	return response
}
