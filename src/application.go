package main

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/spf13/afero"
)

type application struct {
	usage    string
	fsUtil   *afero.Afero
	db       *database
	cmds     map[string]command
	cliArgs  []string
	out      io.Writer
	in       io.Reader
	hasInput bool
	cb       clipboardWriter
}

func newApp(usage string, fsUtil *afero.Afero, db *database, input io.Reader, hasInput bool, output io.Writer, cb *clipboard) *application {
	return &application{
		usage:    usage,
		fsUtil:   fsUtil,
		db:       db,
		out:      output,
		in:       input,
		hasInput: hasInput,
		cb:       cb,
	}
}

func (app *application) registerCommands(cmds []command) {
	app.cmds = make(map[string]command)

	for _, cmd := range cmds {
		app.cmds[cmd.Name] = cmd
	}
}

func (app *application) run(cliArgs []string) bool {
	var (
		cmdArg       string
		helpRequired bool
	)

	app.cliArgs = cliArgs

	if len(cliArgs) > 0 {
		cmdArg = cliArgs[0]

		for _, arg := range cliArgs {
			if arg == "-h" {
				helpRequired = true
			}
		}
	}

	if cmd, ok := app.cmds[cmdArg]; ok {
		if helpRequired {
			app.output(cmd.Usage)
		} else {
			return cmd.Action(app)
		}
	} else {
		app.output(app.usage)
	}

	return true
}

func (app *application) outputSuccess(format string, a ...interface{}) {
	app.outputColor(color.FgGreen, format, a...)
}

func (app *application) outputInfo(format string, a ...interface{}) {
	app.outputColor(color.FgBlue, format, a...)
}

func (app *application) outputError(format string, a ...interface{}) {
	app.outputColor(color.FgRed, format, a...)
}

func (app *application) output(format string, a ...interface{}) {
	fmt.Fprintf(app.out, format+"\n", a...)
}

func (app *application) outputColor(ca color.Attribute, format string, a ...interface{}) {
	c := color.New(ca).SprintFunc()
	s := c(fmt.Sprintf(format, a...))

	app.output("%s", s)
}
