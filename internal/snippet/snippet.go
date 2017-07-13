package snippet

import "strings"

// Snippet represents a snippet stored in the database
type Snippet struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

// Snippets represents more than 1 snippet
type Snippets []Snippet

func (s Snippets) Len() int {
	return len(s)
}

func (s Snippets) Less(i, j int) bool {
	return strings.ToLower(s[i].Filename) < strings.ToLower(s[j].Filename)
}

func (s Snippets) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
