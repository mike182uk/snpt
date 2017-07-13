package snippet

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSnippetsByFilename(t *testing.T) {
	var (
		snptA = Snippet{ID: "3", Filename: "a"}
		snptB = Snippet{ID: "2", Filename: "b"}
		snptC = Snippet{ID: "1", Filename: "c"}
		snpts = Snippets{snptC, snptA, snptB}
	)

	sort.Sort(snpts)

	assert.Equal(t, snptA, snpts[0])
	assert.Equal(t, snptB, snpts[1])
	assert.Equal(t, snptC, snpts[2])
}

func TestJSONEncoding(t *testing.T) {
	snpt := Snippet{
		ID:          "foo",
		Filename:    "bar",
		Description: "baz",
		Content:     "qux",
	}
	expected := []byte(`{"id":"foo","filename":"bar","description":"baz","content":"qux"}`)
	result, err := json.Marshal(snpt)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, result)
}
