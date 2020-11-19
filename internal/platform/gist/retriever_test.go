package gist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func TestRetrieveAll(t *testing.T) {
	retriever := NewRetriever("")
	serverURL := startServer()
	defer stopServer()

	retriever.ghClient.BaseURL = serverURL

	mux.HandleFunc("/gists", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"id":"1"},{"id":"2"}]`)
	})

	expected := []*github.Gist{
		{ID: github.String("1")},
		{ID: github.String("2")},
	}
	result, err := retriever.RetrieveAll()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestRetrieveAllErr(t *testing.T) {
	retriever := NewRetriever("")
	serverURL := startServer()
	defer stopServer()

	retriever.ghClient.BaseURL = serverURL

	mux.HandleFunc("/gists", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	_, err := retriever.RetrieveAll()

	assert.Error(t, err)
}

func TestRetrieveGistFileContent(t *testing.T) {
	retriever := NewRetriever("")
	serverURL := startServer()
	defer stopServer()

	gistFileContent := "foo bar baz"

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, gistFileContent)
	})

	gistFile := github.GistFile{
		RawURL: github.String(serverURL.String()),
	}

	result, err := retriever.RetrieveGistFileContent(&gistFile)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gistFileContent, result)
}

func TestRetrieveGistFileContentErr(t *testing.T) {
	retriever := NewRetriever("")
	serverURL := startServer()
	defer stopServer()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	gistFile := github.GistFile{
		RawURL: github.String(serverURL.String()),
	}

	_, err := retriever.RetrieveGistFileContent(&gistFile)

	assert.EqualError(t, err, fmt.Sprintf("Failed to retrieve gist file content from: %s", serverURL.String()))
}

func startServer() *url.URL {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	url, err := url.Parse(server.URL + "/")

	if err != nil {
		panic(err)
	}

	return url
}

func stopServer() {
	server.Close()
}
