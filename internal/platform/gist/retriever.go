package gist

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/oauth2"
)

// Retriever retrieves Gists and related data from GitHub
type Retriever struct {
	ghClient   *github.Client
	httpClient *gorequest.SuperAgent
}

// NewRetriever returns a new Retriever instance
func NewRetriever(t string) *Retriever {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t},
	)

	c := oauth2.NewClient(context.TODO(), ts)

	return &Retriever{
		ghClient:   github.NewClient(c),
		httpClient: gorequest.New(),
	}
}

// RetrieveAll retrieves all of the Gists for the authenticated user
func (r *Retriever) RetrieveAll() ([]*github.Gist, error) {
	opts := &github.GistListOptions{}

	gists, _, err := r.ghClient.Gists.List(context.TODO(), "", opts)

	if err != nil {
		return nil, err
	}

	return gists, nil
}

// RetrieveGistFileContent retrieves the contents of a file associated with a Gist
func (r *Retriever) RetrieveGistFileContent(file *github.GistFile) (string, error) {
	url := *file.RawURL

	res, content, errs := r.httpClient.Get(url).End()

	if errs != nil || res.StatusCode != 200 {
		return "", fmt.Errorf("Failed to retrieve gist file content from: %s", url)
	}

	return content, nil
}
