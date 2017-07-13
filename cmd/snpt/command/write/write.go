package write

import (
	"errors"
	"fmt"
	"io"

	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/mike182uk/snpt/internal/snippet"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// New returns a new instance of the write command
func New(out io.Writer, in io.Reader, hasInput bool, snptStore *snippet.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "write [snippet ID]",
		Short: "Write a snippet to disk",
		Long: `
Write a snippet to disk.

If you do not provide a snippet ID then a prompt will be shown and you can select the snippet you want to write to disk.

Snpt will read from stdin if provided and attempt to extract a snippet ID from it. The stdin should be formatted like:

	Some random string [snippet ID]

Snpt will parse anything in the square brackets that appears at the end of the string. This is useful for piping into snpt:

	echo 'foo - bar baz [aff9aa71ead70963p3bfa4e49b18d27539f9d9d8]' | snpt write`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fs := afero.NewOsFs()

			snpt, err := cliHelper.ResolveSnippet(args, hasInput, in, snptStore)

			if err != nil || snpt.ID == "" {
				return errors.New("Failed to retrieve snippet from database")
			}

			if err := afero.WriteFile(fs, snpt.Filename, []byte(snpt.Content), 0644); err != nil {
				return fmt.Errorf("Failed to create %s", snpt.Filename)
			}

			cliHelper.PrintSuccess(out, "%s created", snpt.Filename)

			return nil
		},
	}
}
