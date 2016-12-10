package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	name := "foo"
	usage := "bar"

	actionCalled := false

	action := func(app *application) bool {
		actionCalled = true

		return true
	}

	cmd := newCommand(name, usage, action)

	assert.Equal(t, cmd.Name, name)
	assert.Equal(t, cmd.Usage, usage)

	cmd.Action(&application{})

	assert.Equal(t, actionCalled, true)
}
