package main

import "encoding/json"

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

func snippetFromString(s string) (snippet, error) {
	snpt := snippet{}

	if err := json.Unmarshal([]byte(s), &snpt); err != nil {
		return snpt, err
	}

	return snpt, nil
}
