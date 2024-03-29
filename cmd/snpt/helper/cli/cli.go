package cli

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"

	"github.com/fatih/color"
	"github.com/mike182uk/snpt/internal/pb"
	"github.com/mike182uk/snpt/internal/snippet"
	prompt "github.com/segmentio/go-prompt"
)

// GenerateSnippetDescription generates a snippet description
func GenerateSnippetDescription(snpt *pb.Snippet) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	desc := yellow(snpt.Filename)

	if snpt.Description != "" {
		desc += fmt.Sprintf(" - %s", snpt.Description)
	}

	desc += fmt.Sprintf(" [%s]", snpt.GetId())

	return desc
}

// ResolveSnippet resolves a snippet via either a passed argument, standard input or user prompt
func ResolveSnippet(args []string, hasInput bool, in io.Reader, snptStore *snippet.Store) (*pb.Snippet, error) {
	var (
		snpt       pb.Snippet
		snptID     string
		promptOpts []string
		snpts      snippet.Snippets
	)

	snpts, err := snptStore.GetAll()

	if err != nil {
		return &snpt, err
	}

	sort.Sort(snpts)

	// If there are no snippets, return early
	if len(snpts) == 0 {
		return &snpt, nil
	}

	// If there is input data, try and read from it
	if hasInput {
		input, err := io.ReadAll(in)

		if err != nil {
			return &snpt, errors.New("Failed to read from stdin")
		}

		snptID = extractSnippetID(string(input))

		// If there was a second argument passed on the cli, try and use it as an id
	} else if len(args) == 1 {
		snptID = args[0]

		// If there was no input and the id was not passed as an argument, prompt the
		// user to select a snippet
	} else {
		for _, snpt := range snpts {
			promptOpts = append(promptOpts, GenerateSnippetDescription(snpt))
		}

		i := prompt.Choose("Choose a snippet", promptOpts)

		snptID = extractSnippetID(promptOpts[i])
	}

	// Return found snippet
	for _, snpt := range snpts {
		if snpt.GetId() == snptID {
			return snpt, nil
		}

		if snpt.GetFilename() == snptID {
			return snpt, nil
		}
	}

	// Return an empty snippet if no snippet was found
	return &snpt, nil
}

func extractSnippetID(s string) string {
	re := regexp.MustCompile(`(?m:\[([A-Za-z0-9]+)\])`)
	m := re.FindStringSubmatch(s)

	if len(m) == 2 {
		return m[1]
	}

	return ""
}

// PrintSuccess prints a success message
func PrintSuccess(out io.Writer, format string, a ...interface{}) {
	printWithColor(out, color.FgGreen, format, a...)
}

// PrintError prints an error message
func PrintError(out io.Writer, format string, a ...interface{}) {
	printWithColor(out, color.FgRed, format, a...)
}

// PrintInfo prints an info message
func PrintInfo(out io.Writer, format string, a ...interface{}) {
	printWithColor(out, color.FgBlue, format, a...)
}

func printWithColor(out io.Writer, ca color.Attribute, format string, a ...interface{}) {
	c := color.New(ca).SprintFunc()
	s := c(fmt.Sprintf(format, a...))

	fmt.Fprintln(out, s)
}
