package mwdict

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/fsufitch/tagioalisi-bot/merriam-webster/types"
)

var dictionaryAPIBaseURL *url.URL

func init() {
	var err error
	if dictionaryAPIBaseURL, err = url.Parse("https://www.dictionaryapi.com/api/v3/references/collegiate/json"); err != nil {
		panic(err)
	}
}

// Client describes access to the Merriam-Webster dictionary at dictionaryapi.com
type Client interface {
	SearchCollegiate(word string) ([]types.CollegiateResult, []string, error)
	GetInitError() error
}

// BasicClient is a basic HTTP-based client for querying the M-W dictionary
type BasicClient struct {
	APIKey          string
	BaseURL         *url.URL
	UserAgent       string
	Client          *http.Client
	InitFailedError error
}

// NewBasicClient creates a client based on a given API key
func NewBasicClient(apiKey string, userAgent string) *BasicClient {
	var err error = nil
	if apiKey == "" {
		err = errors.New("no merriam-webster api key found")
	}
	return &BasicClient{
		APIKey:          apiKey,
		BaseURL:         dictionaryAPIBaseURL,
		UserAgent:       userAgent,
		Client:          http.DefaultClient,
		InitFailedError: err,
	}
}

func (bc BasicClient) GetInitError() error {
	return bc.InitFailedError
}

// SearchCollegiate implements a search of the collegiate dictionary
func (bc BasicClient) SearchCollegiate(word string) ([]types.CollegiateResult, []string, error) {
	initError := bc.GetInitError()
	if initError != nil {
		return nil, nil, initError
	}

	word = strings.TrimSpace(strings.ToLower(word))

	queryURL := *bc.BaseURL
	queryURL.Path = path.Join(queryURL.Path, word)

	q, _ := url.ParseQuery(queryURL.RawQuery)
	q.Add("key", bc.APIKey)

	queryURL.RawQuery = q.Encode()

	timeoutcontext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutcontext, http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := bc.Client.Do(req)

	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("non-zero status %d; body: %s", response.StatusCode, string(body))
	}

	result := []types.CollegiateResult{}
	suggestions := []string{}

	var err1, err2 error
	if err1 = json.Unmarshal(body, &result); err1 == nil {
		return result, nil, nil
	} else if err2 = json.Unmarshal(body, &suggestions); err2 == nil {
		return nil, suggestions, nil
	}

	return nil, nil, fmt.Errorf("%v %v", err1, err2)
}
