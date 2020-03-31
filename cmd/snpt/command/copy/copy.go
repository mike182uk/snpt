package copy

import (
	"errors"
	"fmt"
	"io"

	"github.com/atotto/clipboard"
	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/mike182uk/snpt/internal/snippet"
	"github.com/spf13/cobra"
)

// New returns a new copy command
func New(out io.Writer, in io.Reader, hasInput bool, snptStore *snippet.Store) *cobra.Command {
	return &cobra.Command{
		Use:     "copy [snippet ID | snippet name]",
		Aliases: []string{"cp"},
		Short:   "Copy a snippet to the clipboard",
		Long: `
Copy a snippet to the clipboard.

If you do not provide a snippet ID or snippet name then a prompt will be shown and you can select the snippet you want to copy to the clipboard.

Snpt will read from stdin if provided and attempt to extract a snippet ID from it. The stdin should be formatted like:

	Some random string [snippet ID]

Snpt will parse anything in the square brackets that appears at the end of the string. This is useful for piping into Snpt:

	echo 'foo - bar baz [aff9aa71ead70963p3bfa4e49b18d27539f9d9d8]' | snpt cp`,
		RunE: func(cmd *cobra.Command, args []string) error {
			snpt, err := cliHelper.ResolveSnippet(args, hasInput, in, snptStore)

			if err != nil || snpt.GetId() == "" {
				return errors.New("Failed to retrieve snippet from database")
			}

			if err := clipboard.WriteAll(snpt.Content); err != nil {
				return fmt.Errorf("Failed to copy %s to the clipboard", snpt.GetFilename())
			}

			cliHelper.PrintSuccess(out, "%s copied to the clipboard", snpt.GetFilename())

			return nil
		},
	}
}
