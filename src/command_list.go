package main

import "sort"

var listCommandName = "ls"

var listCommandUsage = "\nUsage: snpt ls"

var listCommandAction = func(app *application) bool {
	var snpts snippets

	snptStrs, err := app.db.getAll(snippetsBucketName)

	if err != nil {
		app.outputError("Failed to retrieve snippets from database")

		return false
	}

	for _, v := range snptStrs {
		s, _ := snippetFromString(v)

		snpts = append(snpts, s)
	}

	if len(snpts) == 0 {
		app.outputInfo("There are no snippets to list")
	}

	sort.Sort(snpts)

	for _, snpt := range snpts {
		app.output(generateSnippetDescription(snpt))
	}

	return true
}

func newListCommand() command {
	return newCommand(listCommandName, listCommandUsage, listCommandAction)
}
