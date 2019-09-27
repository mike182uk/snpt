package list

import (
	"errors"
	"fmt"
	"io"
	"sort"

	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/mike182uk/snpt/internal/snippet"
	"github.com/spf13/cobra"
)

// New returns a new list command
func New(out io.Writer, snptStore *snippet.Store) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all snippets",
		RunE: func(cmd *cobra.Command, args []string) error {
			snpts, err := snptStore.GetAll()

			if err != nil {
				return errors.New("Failed to retrieve snippets from database")
			}

			if len(snpts) == 0 {
				fmt.Fprintln(out, "There are no snippets to list")

				return nil
			}

			sort.Sort(snpts)

			for _, snpt := range snpts {
				fmt.Fprintln(out, cliHelper.GenerateSnippetDescription(snpt))
			}

			return nil
		},
	}
}
