package azure

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// BingNewsSearch encapsulates functionality of Azure news searches
type BingNewsSearch struct {
	Log *log.Logger
	Key       config.AzureNewsSearchAPIKey
	UserAgent config.UserAgent
}

var newsSearchEndpoint url.URL

func init() {
	url, err := url.Parse("https://api.bing.microsoft.com/v7.0/news/search")
	if err != nil {
		panic("could not parse news search endpoint")
	}
	newsSearchEndpoint = *url
}

// Search performs a search
func (a BingNewsSearch) Search(ctx context.Context, query string, maxNum int32) (*NewsAnswer, error) {
	// See: https://github.com/microsoft/bing-search-sdk-for-python/blob/main/samples/rest/BingCustomSearchV7.py
	//      https://learn.microsoft.com/en-us/bing/search-apis/bing-web-search/reference/endpoints
	url := newsSearchEndpoint

	q := url.Query()
	q.Set("q", query)
	if maxNum > 10 || maxNum < 1 {
		return nil, errors.New("article count must be between 1 and 10")
	}
	q.Set("count", fmt.Sprintf("%d", maxNum))
	
	url.RawQuery = q.Encode()

	req, _ := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	req.Header.Set("User-Agent", string(a.UserAgent))
	req.Header.Set("Ocp-Apim-Subscription-Key", string(a.Key))

	a.Log.Debugf("azure news search, request: %+v", req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	byts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	answer := NewsAnswer{}
	err = json.Unmarshal(byts, &answer)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	a.Log.Debugf("azure news search, response (%d): %+v", resp.StatusCode, answer)

	return &answer, nil
}

// NewsAnswer implements https://learn.microsoft.com/en-us/bing/search-apis/bing-news-search/reference/response-objects#newsanswer
type NewsAnswer struct {
	ReadLink string        `json:"readLink"`
	Articles []NewsArticle `json:"value"`
}

// NewsArticle implements https://learn.microsoft.com/en-us/bing/search-apis/bing-news-search/reference/response-objects#newsarticle
type NewsArticle struct {
	Category      string     `json:"category"`
	DatePublished string     `json:"datePublished"`
	Description   string     `json:"description"`
	Headline      bool       `json:"headline"`
	ID            string     `json:"id"`
	Image         Image      `json:"image"`
	Name          string     `json:"name"`
	Providers     []Provider `json:"provider"`
	URL           string     `json:"url"`
}

type Provider struct {
	Type string `json:"_type"`
	Name string `json:"name"`
}

type Image struct {
	Thumbnail Thumbnail `json:"thumbnail"`
}

type Thumbnail struct {
	URL    string `json:"contentUrl"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
