package wikisupport

import (
	"fmt"
)

// Multi is a struct holding configuration for supported wikis
type Multi struct {
	Wikis   map[string]Single
	Default string
}

// Single encapsulates configuration for one supported wiki
type Single struct {
	ID          string
	Name        string
	DefaultLang string
	IconURL     string
	Client      func(lang string) (Client, error)
}

// Client is an interface for querying a wiki
type Client interface {
	Query(pageName string) (QueryResult, error)
}

// QueryResult holds the results of a single page query
type QueryResult struct {
	Found     bool
	Ambiguous bool
	Title     string
	Text      string
	URL       string
	Thumbnail string
}

// DefaultMultiWikiSupport is the default supported wiki value
var DefaultMultiWikiSupport = Multi{
	Default: "w",
	Wikis: map[string]Single{
		"w": {
			ID:          "w",
			Name:        "Wikipedia",
			DefaultLang: "en",
			IconURL:     "https://upload.wikimedia.org/wikipedia/commons/7/75/Wikipedia_mobile_app_logo.png",
			Client: func(lang string) (Client, error) {
				return NewMediaWikiClient(fmt.Sprintf("https://%s.wikipedia.org/w/api.php", lang))
			},
		},
		"d": {
			ID:          "d",
			Name:        "Wiktionary",
			DefaultLang: "en",
			IconURL:     "https://upload.wikimedia.org/wikipedia/commons/0/07/Wiktsister_en.png",
			Client: func(lang string) (Client, error) {
				return NewMediaWikiClient(fmt.Sprintf("https://%s.wiktionary.org/w/api.php", lang))
			},
		},
		"sw": {
			ID:      "sw",
			Name:    "Wookiepedia",
			IconURL: "https://vignette.wikia.nocookie.net/starwars/images/b/bc/Wiki.png",
			Client: func(lang string) (Client, error) {
				return NewFandomClient("https://starwars.fandom.com/api.php")
			},
		},
		"lotr": {
			ID:      "lotr",
			Name:    "One Wiki to Rule Them All",
			IconURL: "https://vignette.wikia.nocookie.net/lotr/images/b/bc/Wiki.png",
			Client: func(lang string) (Client, error) {
				return NewFandomClient("https://lotr.fandom.com/api.php")
			},
		},
		"tes": {
			ID:      "tes",
			Name:    "The Elder Scrolls Wiki",
			IconURL: "https://i.imgur.com/mgUlAGQ.jpg",
			Client: func(lang string) (Client, error) {
				return NewFandomClient("https://elderscrolls.fandom.com/api.php")
			},
		},
	},
}
