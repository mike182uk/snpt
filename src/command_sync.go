package main

import (
	"regexp"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	"github.com/ararog/timeago"
	"github.com/briandowns/spinner"
	"github.com/google/go-github/github"
	"github.com/parnurzeal/gorequest"
)

const tokenConfigKey = "gh-access-token" // #nosec

var syncCommandName = "sync"

var syncCommandUsage = `
Usage: snpt sync

If you have already supplied a GitHub access token this will be used when
requesting your Gists from GitHub. If you have not supplied one, you will be
prompted to supply one.`

var syncCommandAction = func(app *application) bool {
	token, err := app.db.get(configBucketName + "." + tokenConfigKey)

	if token == "" {
		app.cmds[tokenCommandName].Action(app)

		token, err = app.db.get(configBucketName + "." + tokenConfigKey)
	}

	if err != nil {
		app.outputError("Failed to retrieve GitHub access token from database")

		return false
	}

	if lastSync, _ := getLastSync(app.db); lastSync != "" {
		app.outputInfo("Gists last synced %s", lastSync)
	}

	spinner := getSpinner(app, "Syncing gists...")
	spinner.Start()

	if app.db.empty(snippetsBucketName) != nil {
		spinner.Stop()
		app.outputError("Failed to empty database")

		return false
	}

	ghClient := setupGithubClient(token)
	gists, err := fetchGists(ghClient)

	if err != nil {
		spinner.Stop()
		app.outputError("Failed to retrieve gists")

		return false
	}

	request := gorequest.New()
	gistCount := 0

	for _, gist := range gists {
		for filename, file := range gist.Files {
			var snptStr string
			_, gistContent, errs := request.Get(*file.RawURL).End()

			if errs != nil {
				outputProgressError(spinner, app, "Failed to retrieve gist content from: ", *file.RawURL)

				continue
			}

			gistID := getGistIDFromRawURL(*file.RawURL)

			if gistID == "" {
				outputProgressError(spinner, app, "Failed to extract gist ID from URL: ", *file.RawURL)

				continue
			}

			snpt := snippet{
				ID:          gistID,
				Filename:    string(filename),
				Description: *gist.Description,
				Content:     gistContent,
			}

			snptStr, err = snpt.toString()

			if err != nil {
				outputProgressError(spinner, app, "Failed to convert snippet to string: ", snpt.Filename)

				continue
			}

			err = app.db.set(snippetsBucketName+"."+gistID, snptStr)

			if err != nil {
				outputProgressError(spinner, app, "Failed to save snippet: %s to the database", snpt.Filename)

				continue
			}

			gistCount++
		}
	}

	err = app.db.set(miscBucketName+".last-sync", time.Now().Format(time.RFC3339))

	if err != nil {
		app.outputError("Failed to save last sync time to the database")

		return false
	}

	spinner.Stop()

	app.outputSuccess("%d gist(s) synced", gistCount)

	return true
}

func getSpinner(app *application, suffix string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	s.Suffix = " " + suffix
	s.Writer = app.out

	return s
}

func setupGithubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(context.TODO(), ts)

	client := github.NewClient(tc)

	return client
}

func fetchGists(client *github.Client) ([]*github.Gist, error) {
	opts := &github.GistListOptions{}

	gists, _, err := client.Gists.List("", opts)

	if err != nil {
		return nil, err
	}

	return gists, nil
}

func getGistIDFromRawURL(url string) string {
	re := regexp.MustCompile("/raw/(.*)/")
	m := re.FindStringSubmatch(url)

	if len(m) == 2 {
		return m[1]
	}

	return ""
}

func getLastSync(db *database) (string, error) {
	ls, err := db.get(miscBucketName + ".last-sync")

	if err != nil {
		return "", nil
	}

	if ls != "" {
		t, err := time.Parse(time.RFC3339, ls)

		if err != nil {
			return "", err
		}

		ls, err = timeago.TimeAgoWithTime(time.Now(), t)

		if err != nil {
			return "", err
		}
	}

	return ls, nil
}

func outputProgressError(s *spinner.Spinner, app *application, format string, a ...interface{}) {
	s.Stop()

	app.outputError(format, a...)

	s.Restart()
}

func newSyncCommand() command {
	return newCommand(syncCommandName, syncCommandUsage, syncCommandAction)
}
