package main

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/boltdb/bolt"
	"github.com/spf13/afero"
)

const appDirName = ".snpt"
const dbName = "snpt.db"

const snippetsBucketName = "Snippets"
const configBucketName = "Config"
const miscBucketName = "Misc"

const usage = `
Commands:
    cp       Copy a snippet to the clipboard
    ls       List all available snippets
    sync     Download gists from github and store locally
    write    Write a snippet to disk
    token    Set your GitHub access token

To view usage information for a command, use the -h flag when running it:

  snpt cp -h
`

func main() {
	// setup the app filesystem
	fs := afero.NewOsFs()

	// setup the app directory where app related files will be stored
	appDirPath, err := setupAppDir(fs)

	if err != nil {
		fmt.Println("Failed to setup app directory")
		os.Exit(1)
	}

	// setup the app database
	db, err := setupDB(appDirPath)

	if err != nil {
		fmt.Println("Failed to setup database")

		db.boltDB.Close()

		os.Exit(1)
	}

	// get the cli args
	cliArgs := os.Args[1:]

	// is there any input on stdin?
	hasInput := false
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		hasInput = true
	}

	// intialise the application
	app := newApp(usage, &afero.Afero{Fs: fs}, db, os.Stdin, hasInput, os.Stdout, &clipboard{})

	cmds := []command{
		newTokenCommand(),
		newListCommand(),
		newSyncCommand(),
		newCopyCommand(),
		newWriteCommand(),
	}

	app.registerCommands(cmds)

	// run the app
	ok := app.run(cliArgs)

	// close the database
	db.boltDB.Close()

	// exit with the correct code
	if ok == true {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func setupAppDir(fs afero.Fs) (string, error) {
	u, err := user.Current()

	if err != nil {
		return "", err
	}

	appDirPath := path.Join(u.HomeDir, appDirName)

	if err := fs.MkdirAll(appDirPath, 0644); err != nil {
		return "", err
	}

	return appDirPath, nil
}

func setupDB(appDirPath string) (*database, error) {
	dbPath := path.Join(appDirPath, dbName)

	boltDB, err := bolt.Open(dbPath, 0644, nil)

	if err != nil {
		return nil, err
	}

	db := newDatabase(boltDB)
	bucketNames := []string{
		configBucketName,
		snippetsBucketName,
		miscBucketName,
	}

	if err = db.init(bucketNames); err != nil {
		return nil, err
	}

	return db, nil
}
