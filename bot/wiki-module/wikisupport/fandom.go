package wikisupport

import (
	"encoding/json"
	"fmt"
	"strings"

	"cgt.name/pkg/go-mwclient"
	"github.com/PuerkitoBio/goquery"
	"github.com/antonholmquist/jason"
	"github.com/pkg/errors"
)

const fandomUserAgent = "TagioalisiBot / mwclient; https://github.com/fsufitch/tagioalisi-bot"

// FandomClient is capable of performing queries to a MediaWiki-compliant API
type FandomClient struct {
	client *mwclient.Client
}

// NewFandomClient creates a new FandomClient
func NewFandomClient(url string) (*FandomClient, error) {
	client, err := mwclient.New(url, mediaWikiUserAgent)
	if err != nil {
		return nil, err
	}
	return &FandomClient{client: client}, nil
}

// Query performs a wiki query
func (c FandomClient) Query(title string) (QueryResult, error) {
	if strings.Contains(title, "|") {
		return QueryResult{}, errors.New("title may not contain a pipe (|)")
	}
	jasonResult, err := c.client.Get(map[string]string{
		"action":    "query",
		"prop":      "info|pageprops|revisions|images",
		"titles":    title,
		"inprop":    "url|displaytitle",
		"redirects": "",
		"rvprop":    "content",
		"rvlimit":   "1",
		"rvparse":   "",
	})
	if err != nil {
		return QueryResult{}, errors.Wrap(err, "error executing query")
	}

	jsonResult, err := reUnmarshalJasonFandom(jasonResult)

	if len(jsonResult.Query.Pages) < 1 {
		return QueryResult{Found: false}, nil
	}

	var pageID string
	for k := range jsonResult.Query.Pages {
		pageID = k
		break
	}

	if pageID == "" || jsonResult.Query.Pages[pageID].ID == 0 {
		return QueryResult{Found: false}, nil
	}
	resultPage := jsonResult.Query.Pages[pageID]

	text := ""
	if len(resultPage.Revisions) > 0 {
		html := resultPage.Revisions[0].Star
		html = "<toplevel>" + html + "</toplevel>"

		gq, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			text = fmt.Sprintf("Error extracting text: %v", err)
		} else if gq.Find("toplevel > p").Length() == 0 {
			text = fmt.Sprintf("Error extracting text: could not find first paragraph")
		} else {
			text = gq.Find("toplevel > p").First().Text()
		}
	}

	result := QueryResult{
		Found:     true,
		Ambiguous: false,
		Title:     resultPage.Title,
		URL:       resultPage.URL,
		Text:      text,
		//Thumbnail: resultPage.Thumbnail.URL,
	}

	if strings.Contains(resultPage.Title, "(disambiguation)") {
		result.Ambiguous = true
	}

	return result, nil
}

func reUnmarshalJasonFandom(jObj *jason.Object) (fandomQueryResult, error) {
	data, err := jObj.Marshal()
	if err != nil {
		return fandomQueryResult{}, err
	}
	obj := fandomQueryResult{}
	err = json.Unmarshal(data, &obj)
	return obj, err
}

type fandomQueryResult struct {
	Query struct {
		Pages map[string]struct {
			ID       int    `json:"pageid"`
			Title    string `json:"displaytitle"`
			URL      string `json:"fullurl"`
			Text     string `json:"extract"`
			Redirect bool   `json:"redirect"`
			Images   []struct {
				Title string `json:"title"`
			} `json:"thumbnail"`
			Props     map[string]interface{} `json:"pageprops"`
			Revisions []struct {
				Star string `json:"*"`
			}
		} `json:"pages"`
	} `json:"query"`
}
