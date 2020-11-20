package snippet

import (
	"strings"

	"github.com/mike182uk/snpt/internal/pb"
)

// Snippets represents more than 1 snippet
type Snippets []*pb.Snippet

func (s Snippets) Len() int {
	return len(s)
}

func (s Snippets) Less(i, j int) bool {
	return strings.ToLower(s[i].GetFilename()) < strings.ToLower(s[j].GetFilename())
}

func (s Snippets) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
