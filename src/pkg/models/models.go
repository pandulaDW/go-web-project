package models

import (
	"errors"
	"time"
)

// ErrNoRecord is an error generated when no results are found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is a struct to define the db model
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
