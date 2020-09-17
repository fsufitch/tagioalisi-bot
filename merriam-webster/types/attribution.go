package types

type withQuoteAttribution struct {
	// https://dictionaryapi.com/products/json#sec-2.aq
	QuoteAttribution QuoteAttribution `json:"aq"`
}

// QuoteAttribution https://dictionaryapi.com/products/json#sec-2.aq
type QuoteAttribution struct {
	Author    string `json:"auth"`
	Source    string `json:"source"`
	Date      string `json:"aqdate"`
	SubSource struct {
		Source string `json:"source"`
		Date   string `json:"aqdate"`
	} `json:"subsource"`
}
