package main

import (
	"encoding/json"
	"strings"
)

type snippet struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func (snpt snippet) toString() (string, error) {
	b, err := json.Marshal(snpt)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

type snippets []snippet

func (snpts snippets) Len() int {
	return len(snpts)
}

func (snpts snippets) Less(i, j int) bool {
	return strings.ToLower(snpts[i].Filename) < strings.ToLower(snpts[j].Filename)
}

func (snpts snippets) Swap(i, j int) {
	snpts[i], snpts[j] = snpts[j], snpts[i]
}

func snippetFromString(s string) (snippet, error) {
	snpt := snippet{}

	if err := json.Unmarshal([]byte(s), &snpt); err != nil {
		return snpt, err
	}

	return snpt, nil
}
