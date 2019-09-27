package print

import (
	"errors"
	"io"

	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/mike182uk/snpt/internal/snippet"
	"github.com/spf13/cobra"
)

// New returns a new print command
func New(out io.Writer, in io.Reader, hasInput bool, snptStore *snippet.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "print [snippet ID]",
		Short: "Print a snippet to the screen",
		Long: `
Print a snippet to the screen.

If you do not provide a snippet ID then a prompt will be shown and you can select the snippet you want to print to the screen.

Snpt will read from stdin if provided and attempt to extract a snippet ID from it. The stdin should be formatted like:

	Some random string [snippet ID]

Snpt will parse anything in the square brackets that appears at the end of the string. This is useful for piping into snpt:

	echo 'foo - bar baz [aff9aa71ead70963p3bfa4e49b18d27539f9d9d8]' | snpt print`,
		RunE: func(cmd *cobra.Command, args []string) error {
			snpt, err := cliHelper.ResolveSnippet(args, hasInput, in, snptStore)

			if err != nil || snpt.GetId() == "" {
				return errors.New("Failed to retrieve snippet from database")
			}

			cliHelper.PrintInfo(out, "\n%s", snpt.Content)

			return nil
		},
	}
}
