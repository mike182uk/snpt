package main

import (
	"testing"

	clipboardMgr "github.com/atotto/clipboard"
	"github.com/stretchr/testify/assert"
)

func TestWriteToClipboard(t *testing.T) {
	cb := new(clipboard)
	current, _ := clipboardMgr.ReadAll()

	err := cb.writeToClipboard("foo bar baz")

	if err != nil {
		t.Fatalf("Failed to write to the clipboard: %s", err)
	}

	text, err := clipboardMgr.ReadAll()

	if err != nil {
		t.Fatalf("Failed to read from the clipboard: %s", err)
	}

	assert.Equal(t, text, "foo bar baz")

	// restore clipboard
	clipboardMgr.WriteAll(current)
}
