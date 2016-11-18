package main

import "github.com/segmentio/go-prompt"

var tokenCommandName = "token"

var tokenCommandUsage = `
Usage: snpt token

You will be prompted to supply your GitHub access token. This will be saved
in Snpt's database for use when retrieving Gists from GitHub.`

var tokenCommandAction = func(app *application) bool {
	token := prompt.PasswordMasked("Enter your GitHub access token")

	if err := app.db.set(configBucketName+".gh-access-token", token); err != nil {
		app.outputError("Failed to save GitHub access token")

		return false
	}

	app.outputSuccess("GitHub access token saved")

	return true
}

func newTokenCommand() command {
	return newCommand(tokenCommandName, tokenCommandUsage, tokenCommandAction)
}
