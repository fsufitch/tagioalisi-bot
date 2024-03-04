package types

type withMetadata struct {
	Metadata Metadata `json:"meta"`
}

// Metadata https://dictionaryapi.com/products/json#sec-2.meta
type Metadata struct {
	ID        string   `json:"id"`
	UUID      string   `json:"uuid"`
	SortKey   string   `json:"sort"`
	Source    string   `json:"src"`
	Section   string   `json:"section"`
	Stems     []string `json:"stems"`
	Offensive bool     `json:"offensive"`
}
