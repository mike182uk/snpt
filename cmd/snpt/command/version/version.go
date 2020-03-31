package list

import (
	"io"

	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/spf13/cobra"
)

// New returns a new version command
func New(out io.Writer, version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of snpt",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHelper.PrintInfo(out, "snpt v%s", version)

			return nil
		},
	}
}
