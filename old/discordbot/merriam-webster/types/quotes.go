package types

type withQuotes struct {
	// https://dictionaryapi.com/products/json#sec-2.quotes
	Quotes []Quote `json:"quotes"`
}

// Quote https://dictionaryapi.com/products/json#sec-2.quotes
type Quote struct {
	Text             string           `json:"t"`
	QuoteAttribution QuoteAttribution `json:"aq"`
}
