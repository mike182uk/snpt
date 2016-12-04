package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"

	"github.com/fatih/color"
	"github.com/segmentio/go-prompt"
)

func generateSnippetDescription(snpt snippet) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	desc := yellow(snpt.Filename)

	if snpt.Description != "" {
		desc += fmt.Sprintf(" - %s", snpt.Description)
	}

	desc += fmt.Sprintf(" [%s]", snpt.ID)

	return desc
}

func extractSnippetIDFromString(s string) string {
	re := regexp.MustCompile(`\[([A-Za-z0-9]+)\]$`)

	m := re.FindStringSubmatch(s)

	if len(m) == 2 {
		return m[1]
	}

	return ""
}

func resolveSnippet(cliArgs []string, hasInput bool, input io.Reader, db *database) (snippet, error) {
	var (
		snptID     string
		promptOpts []string
		snpts      snippets
	)

	// fetch snippets from db
	snptStrs, err := db.getAll(snippetsBucketName)

	if err != nil {
		return snippet{}, err
	}

	for _, v := range snptStrs {
		s, _ := snippetFromString(v)

		snpts = append(snpts, s)
	}

	sort.Sort(snpts)

	// if there are no snippets, return early
	if len(snpts) == 0 {
		return snippet{}, nil
	}

	// if there is input data, try and read from it
	if hasInput == true {
		reader := bufio.NewReader(input)

		line, _, err := reader.ReadLine()

		if err != nil {
			return snippet{}, errors.New("Failed to read from input")
		}

		snptID = extractSnippetIDFromString(string(line))

		// if there was a second argument passed on the cli, try and use it as an id
	} else if len(cliArgs) == 2 {
		snptID = cliArgs[1]

		// if there was no input and the id was not passed as a cli argument, prompt the
		// user to select a snippet
	} else {
		for _, v := range snpts {
			promptOpts = append(promptOpts, generateSnippetDescription(v))
		}

		i := prompt.Choose("Choose a snippet", promptOpts)

		snptID = extractSnippetIDFromString(promptOpts[i])
	}

	// return found snippet
	for _, snpt := range snpts {
		if snpt.ID == snptID {
			return snpt, nil
		}
	}

	// return an empty snippet if no snippet was found
	return snippet{}, nil
}
