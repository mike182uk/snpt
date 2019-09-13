package snippet

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSnippetsByFilename(t *testing.T) {
	var (
		snptA = Snippet{Id: "3", Filename: "a"}
		snptB = Snippet{Id: "2", Filename: "b"}
		snptC = Snippet{Id: "1", Filename: "c"}
		snpts = Snippets{snptC, snptA, snptB}
	)

	sort.Sort(snpts)

	assert.Equal(t, snptA, snpts[0])
	assert.Equal(t, snptB, snpts[1])
	assert.Equal(t, snptC, snpts[2])
}
