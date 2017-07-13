package token

import (
	"errors"
	"io"

	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"
	"github.com/mike182uk/snpt/internal/config"
	prompt "github.com/segmentio/go-prompt"
	"github.com/spf13/cobra"
)

// New returns a new instance of the token command
func New(out io.Writer, c *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "token",
		Short: "Set your GitHub access token",
		Long: `
Set your GitHub access token.

You will be prompted to supply your GitHub access token. This will be saved in snpt's database for use when retrieving your Gists from GitHub.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			token := prompt.PasswordMasked("Enter your GitHub access token")

			// TODO: if user enters an empty string, re-prompt them for a token, if they cancel (ctrl + c) do nothing
			if token == "" {
				cliHelper.PrintError(out, "GitHub access token was not supplied")

				return nil
			}

			if err := c.Set(config.TokenKey, token); err != nil {
				return errors.New("Failed to save GitHub access token")
			}

			cliHelper.PrintSuccess(out, "GitHub access token saved")

			return nil
		},
	}
}
