package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	snpt := snippet{
		ID:          "foo123",
		Filename:    "foo.txt",
		Description: "foo text file",
		Content:     "foo bar baz",
	}

	s, err := snpt.toString()

	if err != nil {
		t.Fatalf("Failed to convert snippet to string: %s", err)
	}

	assert.Equal(t, s, `{"id":"foo123","filename":"foo.txt","description":"foo text file","content":"foo bar baz"}`)
}

func TestSnippetFromString(t *testing.T) {
	snptStr := `{"id":"foo123","filename":"foo.txt","description":"foo text file","content":"foo bar baz"}`

	snpt, err := snippetFromString(snptStr)

	if err != nil {
		t.Fatalf("Failed to create snippet from string: %s", err)
	}

	assert.Equal(t, snpt, snippet{
		ID:          "foo123",
		Filename:    "foo.txt",
		Description: "foo text file",
		Content:     "foo bar baz",
	})
}
