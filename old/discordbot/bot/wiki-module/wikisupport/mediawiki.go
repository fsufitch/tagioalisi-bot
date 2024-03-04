package wikisupport

import (
	"encoding/json"
	"strings"

	"cgt.name/pkg/go-mwclient"
	"github.com/antonholmquist/jason"
	"github.com/pkg/errors"
)

const mediaWikiUserAgent = "TagioalisiBot / mwclient; https://github.com/fsufitch/tagioalisi-bot"

// MediaWikiClient is capable of performing queries to a MediaWiki-compliant API
type MediaWikiClient struct {
	client *mwclient.Client
}

// NewMediaWikiClient creates a new MediaWikiClient
func NewMediaWikiClient(url string) (*MediaWikiClient, error) {
	client, err := mwclient.New(url, mediaWikiUserAgent)
	if err != nil {
		return nil, err
	}
	return &MediaWikiClient{client: client}, nil
}

// Query performs a wiki query
func (c MediaWikiClient) Query(title string) (QueryResult, error) {
	if strings.Contains(title, "|") {
		return QueryResult{}, errors.New("title may not contain a pipe (|)")
	}
	jasonResult, err := c.client.Get(map[string]string{
		"action":      "query",
		"prop":        "extracts|info|pageimages|pageprops",
		"titles":      title,
		"exlimit":     "1",
		"exsentences": "5",
		"explaintext": "",
		"piprop":      "thumbnail",
		"pithumbsize": "256",
		"inprop":      "url|displaytitle",
		"redirects":   "",
	})
	if err != nil {
		return QueryResult{}, errors.Wrap(err, "error executing query")
	}

	jsonResult, err := reUnmarshalJasonMediaWiki(jasonResult)

	if len(jsonResult.Query.Pages) < 1 {
		return QueryResult{Found: false}, nil
	}

	jsonPage := jsonResult.Query.Pages[0]
	if jsonPage.ID == 0 {
		return QueryResult{Found: false}, nil
	}

	result := QueryResult{
		Found:     true,
		Ambiguous: false,
		Title:     jsonPage.Title,
		URL:       jsonPage.URL,
		Text:      jsonPage.Text,
		Thumbnail: jsonPage.Thumbnail.URL,
	}

	if _, ok := jsonPage.Props["disambiguation"]; ok {
		result.Ambiguous = true
	}

	return result, nil
}

func reUnmarshalJasonMediaWiki(jObj *jason.Object) (mediaWikiQueryResult, error) {
	data, err := jObj.Marshal()
	if err != nil {
		return mediaWikiQueryResult{}, err
	}
	println(string(data))
	obj := mediaWikiQueryResult{}
	err = json.Unmarshal(data, &obj)
	return obj, err
}

type mediaWikiQueryResult struct {
	Query struct {
		Pages []struct {
			ID        int    `json:"pageid"`
			Title     string `json:"displaytitle"`
			URL       string `json:"fullurl"`
			Text      string `json:"extract"`
			Redirect  bool   `json:"redirect"`
			Thumbnail struct {
				URL string `json:"source"`
			} `json:"thumbnail"`
			Props map[string]interface{} `json:"pageprops"`
		} `json:"pages"`
	} `json:"query"`
}
