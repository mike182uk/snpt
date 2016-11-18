package main

var copyCommandName = "cp"

var copyCommandUsage = `
Usage: snpt cp [snippetId]

If you do not provide a snippetId then a prompt will be shown and you can
select the snippet you want to copy to the clipboard.

Snpt will read from stdin if provided and attempt to extract a snippet ID
from it. The stdin should be formatted like:

  Some random string [snippetId]

Snpt will parse anything in the sqaure brackets that appears at the end of
the string. This is useful for piping into snpt:

  echo 'foo - bar baz [aff9aa71ead70963p3bfa4e49b18d27539f9d9d8]' | snpt cp`

var copyCommandAction = func(app *application) bool {
	snpt, err := resolveSnippet(app.cliArgs, app.hasInput, app.in, app.db)

	if err != nil {
		app.outputError("Failed to search for snippet")

		return false
	}

	if (snpt == snippet{}) {
		app.outputError("Snippet not found")

		return false
	}

	if err := app.cb.writeToClipboard(snpt.Content); err != nil {
		app.outputError("Failed to create %s", snpt.Filename)

		return false
	}

	app.outputSuccess("%s copied to the clipboard", snpt.Filename)

	return true
}

func newCopyCommand() command {
	return newCommand(copyCommandName, copyCommandUsage, copyCommandAction)
}
