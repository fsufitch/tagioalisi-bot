package types

type withHeadwordInfo struct {
	HeadwordInfo HeadwordInfo `json:"hwi"`
}

// HeadwordInfo https://dictionaryapi.com/products/json#sec-2.hwi
type HeadwordInfo struct {
	Headword string `json:"hw"`
	withPronounciations
}
