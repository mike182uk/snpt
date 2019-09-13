package list

import (
	"io"

	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/spf13/cobra"
)

const version = "3.0.0"

// New returns a new instance of the version command
func New(out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of snpt",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHelper.PrintInfo(out, "snpt v%s", version)

			return nil
		},
	}
}
