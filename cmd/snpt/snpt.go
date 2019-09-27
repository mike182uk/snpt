package main

import (
	"os"
	"os/user"
	"path"

	copyCmd "github.com/mike182uk/snpt/cmd/snpt/command/copy"
	listCmd "github.com/mike182uk/snpt/cmd/snpt/command/list"
	printCmd "github.com/mike182uk/snpt/cmd/snpt/command/print"
	syncCmd "github.com/mike182uk/snpt/cmd/snpt/command/sync"
	tokenCmd "github.com/mike182uk/snpt/cmd/snpt/command/token"
	versionCmd "github.com/mike182uk/snpt/cmd/snpt/command/version"
	writeCmd "github.com/mike182uk/snpt/cmd/snpt/command/write"
	cliHelper "github.com/mike182uk/snpt/cmd/snpt/helper/cli"

	"github.com/mike182uk/snpt/internal/config"
	"github.com/mike182uk/snpt/internal/platform/storage"
	"github.com/mike182uk/snpt/internal/snippet"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const (
	appDir = ".snpt"
	dbName = "snpt.db"
)

func main() {
	rootCmd := cobra.Command{
		Use:           "snpt",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	in := os.Stdin
	out := rootCmd.OutOrStdout()

	// setup storage
	storage, err := initStorage()

	if err != nil {
		cliHelper.PrintError(out, "Failed to open database")

		os.Exit(1)
	}

	// setup snippet store
	snptStore, err := snippet.NewStore(storage)

	if err != nil {
		cliHelper.PrintError(out, "Failed to initialise snippet store")

		os.Exit(1)
	}

	// setup config
	config, err := config.New(storage)

	if err != nil {
		cliHelper.PrintError(out, "Failed to initialise config")

		os.Exit(1)
	}

	// is there any input on stdin?
	hasInput := false
	stat, _ := in.Stat()

	if stat.Mode()&os.ModeNamedPipe != 0 {
		hasInput = true
	}

	// register commands
	rootCmd.AddCommand(
		copyCmd.New(out, in, hasInput, snptStore),
		listCmd.New(out, snptStore),
		printCmd.New(out, in, hasInput, snptStore),
		syncCmd.New(out, config, snptStore),
		tokenCmd.New(out, config),
		versionCmd.New(out),
		writeCmd.New(out, in, hasInput, snptStore),
	)

	// run the app
	if err := rootCmd.Execute(); err != nil {
		cliHelper.PrintError(out, err.Error())

		os.Exit(1)
	}
}

func initStorage() (storage.BucketKeyValueStore, error) {
	storePath, err := getStorePath()

	if err != nil {
		return nil, err
	}

	s, err := storage.NewBoltStore(storePath)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func getStorePath() (string, error) {
	fs := afero.NewOsFs()
	u, err := user.Current()

	if err != nil {
		return "", err
	}

	dir := path.Join(u.HomeDir, appDir)

	if err := fs.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return path.Join(dir, dbName), nil
}
