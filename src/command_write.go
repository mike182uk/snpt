package main

var writeCommandName = "write"

var writeCommandUsage = `
Usage: snpt write [snippetId]

If you do not provide a snippetId then a prompt will be shown and you can
select the snippet you want to write to disk.

Snpt will read from stdin if provided and attempt to extract a snippet ID
from it. The stdin should be formatted like:

  Some random string [snippetId]

Snpt will parse anything in the sqaure brackets that appears at the end of
the string. This is useful for piping into snpt:

  echo 'foo - bar [aff9aa71ead70963p3bfa4e49b18d27539f9d9d8]' | snpt write`

var writeCommandAction = func(app *application) bool {
	snpt, err := resolveSnippet(app.cliArgs, app.hasInput, app.in, app.db)

	if err != nil {
		app.outputError("Failed to search for snippet")

		return false
	}

	if (snpt == snippet{}) {
		app.outputError("Snippet not found")

		return false
	}

	if err = app.fsUtil.WriteFile(snpt.Filename, []byte(snpt.Content), 0644); err != nil {
		app.outputError("Failed to create %s", snpt.Filename)

		return false
	}

	app.outputSuccess("%s created", snpt.Filename)

	return true
}

func newWriteCommand() command {
	return newCommand(writeCommandName, writeCommandUsage, writeCommandAction)
}
