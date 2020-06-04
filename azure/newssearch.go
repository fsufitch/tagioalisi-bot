package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/newssearch"
	"github.com/Azure/go-autorest/autorest"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/pkg/errors"
)

// OnlineNewsSearch encapsulates functionality of Azure news searches
type OnlineNewsSearch struct {
	client *newssearch.NewsClient
}

// ProvideOnlineNewsSearch creates a NewsSearch based on the available API key
func ProvideOnlineNewsSearch(key config.AzureNewsSearchAPIKey) *OnlineNewsSearch {
	ans := OnlineNewsSearch{}
	if key != "" {
		client := newssearch.NewNewsClient()
		client.Authorizer = autorest.NewCognitiveServicesAuthorizer(string(key))
		client.AddToUserAgent("tagioalisi-bot")
		ans.client = &client
	}
	return &ans
}

// Ready returns whether the NewsSearch is ready to use
func (a OnlineNewsSearch) Ready() bool {
	return a.client != nil
}

// Search performs a search
func (a OnlineNewsSearch) Search(ctx context.Context, query string, maxNum int32) (NewsResults, error) {
	if maxNum < 1 || maxNum > 10 {
		return nil, errors.New("only 1-10 return articles supported")
	}

	news, err := a.client.Search(
		context.Background(), // context
		query,                // query keyword
		"",                   // Accept-Language header
		"",                   // User-Agent header
		"",                   // X-MSEdge-ClientID header
		"",                   // X-MSEdge-ClientIP header
		"",                   // X-Search-Location header
		"",                   // country code
		&maxNum,              // count
		newssearch.Month,     // freshness
		"",                   // market
		nil,                  // offset
		nil,                  // original image
		newssearch.Strict,    // safe search
		"",                   // set lang
		"",                   // sort by
		nil,                  // text decorations
		newssearch.Raw,       // text format
	)

	if err != nil {
		return nil, fmt.Errorf("unexpected error querying Azure news search: %w", err)
	}
	if news.Value == nil {
		return nil, errors.New("search response .Value was nil")
	}

	return OnlineNewsResults{news}, nil
}

// OnlineNewsResults contains a result set of a search
type OnlineNewsResults struct {
	news newssearch.News
}

// Len returns the number of articles in the result set
func (r OnlineNewsResults) Len() int {
	return len(*r.news.Value)
}

// Get returns the n'th result in the result set
func (r OnlineNewsResults) Get(idx int) (NewsResult, bool) {
	if idx >= r.Len() {
		return nil, false
	}

	return OnlineNewsResult{(*r.news.Value)[idx]}, true
}

// OnlineNewsResult is a single article in the result set
type OnlineNewsResult struct {
	article newssearch.NewsArticle
}

// Title returns the article's title
func (r OnlineNewsResult) Title() string {
	if r.article.Name == nil {
		return ""
	}
	return *r.article.Name
}

// Description returns the article's title
func (r OnlineNewsResult) Description() string {
	if r.article.Description == nil {
		return ""
	}
	return *r.article.Description
}

// Source returns the article's first cited provider/source
func (r OnlineNewsResult) Source() string {
	if r.article.Provider == nil || len(*r.article.Provider) == 0 {
		println("source was nil or empty", r.article.Provider)
		return ""
	}

	if p, ok := (*r.article.Provider)[0].AsOrganization(); ok {
		return *p.Name
	}

	if p, ok := (*r.article.Provider)[0].AsThing(); ok {
		return *p.Name
	}
	println("provider was not a thing or org", (*r.article.Provider)[0])
	return ""
}

// ThumbnailURL returns the article's thumbnail image
func (r OnlineNewsResult) ThumbnailURL() string {
	if r.article.Image != nil &&
		r.article.Image.Thumbnail != nil &&
		r.article.Image.Thumbnail.ContentURL != nil {
		return *r.article.Image.Thumbnail.ContentURL
	}
	print("thumbnail was nil")
	return ""
}

// URL returns the URL to the article itself
func (r OnlineNewsResult) URL() string {
	if r.article.URL == nil {
		return ""
	}
	return *r.article.URL
}

// NewsSearch is a general interface for a news search service
type NewsSearch interface {
	Ready() bool
	Search(ctx context.Context, query string, maxNum int32) (NewsResults, error)
}

// NewsResults is a general interface for a collection of news
type NewsResults interface {
	Len() int
	Get(int) (NewsResult, bool)
}

// NewsResult is a general interface for a single news article
type NewsResult interface {
	Title() string
	Description() string
	Source() string
	ThumbnailURL() string
	URL() string
}
