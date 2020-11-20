package sync

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/ararog/timeago"
	"github.com/briandowns/spinner"
	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/mike182uk/snpt/internal/config"
	"github.com/mike182uk/snpt/internal/pb"
	"github.com/mike182uk/snpt/internal/platform/gist"
	"github.com/mike182uk/snpt/internal/snippet"
	"github.com/spf13/cobra"
)

// New returns a new sync command
func New(out io.Writer, c *config.Config, snptStore *snippet.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Sync snippets",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, err := c.Get(config.TokenKey)

			if err != nil {
				return errors.New("Failed to retrieve GitHub access token from database")
			}

			if token == "" {
				cliHelper.PrintInfo(out, "No access token set. Run [snpt token] to set your GitHub access token")

				return nil
			}

			if lastSync, _ := getLastSync(c); lastSync != "" {
				cliHelper.PrintInfo(out, "Gists last synced %s", lastSync)
			}

			spinner := getSpinner(out, "Syncing gists...")
			spinner.Start()
			defer spinner.Stop()

			if snptStore.DeleteAll() != nil {
				spinner.Stop()

				return errors.New("Failed to empty database")
			}

			gistRetriever := gist.NewRetriever(token)
			gists, err := gistRetriever.RetrieveAll()

			if err != nil {
				spinner.Stop()

				return errors.New("Failed to retrieve gists from GitHub")
			}

			gistCount := 0
			ignoreRe := regexp.MustCompile(`\[snpt:ignore\]`)

			for _, gist := range gists {
				if ignoreRe.FindStringIndex(*gist.Description) != nil {
					continue
				}

				for filename, file := range gist.Files {
					fileContent, err := gistRetriever.RetrieveGistFileContent(&file)

					if err != nil {
						printProgressError(spinner, out, err.Error())

						continue
					}

					h := md5.New()
					_, err = h.Write([]byte(*file.RawURL))

					if err != nil {
						printProgressError(spinner, out, "Failed to generate ID for gist: %s", *file.RawURL)

						continue
					}

					gistID := hex.EncodeToString(h.Sum(nil))

					snpt := pb.Snippet{
						Id:          gistID,
						Filename:    string(filename),
						Description: *gist.Description,
						Content:     fileContent,
					}

					err = snptStore.Put(&snpt)

					if err != nil {
						printProgressError(spinner, out, "Failed to save snippet: %s to the database", snpt.Filename)

						continue
					}

					gistCount++
				}
			}

			_ = c.Set(config.LastSyncKey, time.Now().Format(time.RFC3339))

			spinner.Stop()

			cliHelper.PrintSuccess(out, "%d gist(s) synced", gistCount)

			return nil
		},
	}
}

func getSpinner(out io.Writer, suffix string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	s.Suffix = " " + suffix
	s.Writer = out

	return s
}

func getLastSync(c *config.Config) (string, error) {
	ls, err := c.Get(config.LastSyncKey)

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

func printProgressError(s *spinner.Spinner, out io.Writer, format string, a ...interface{}) {
	s.Stop()

	cliHelper.PrintError(out, format, a...)

	s.Restart()
}
