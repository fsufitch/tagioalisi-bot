package types

type withPronounciations struct {
	Pronounciations []Pronounciation `json:"prs"`
}

// Pronounciation https://dictionaryapi.com/products/json#sec-2.prs
type Pronounciation struct {
	MerriamWebsterFormat string `json:"mw"`
	LabelBefore          string `json:"l"`
	LabelAfter           string `json:"l2"`
	Punctuation          string `json:"pun"`
	Sound                struct {
		Filename string `json:"audio"`
	} `json:"sound"`
}

// Sound https://dictionaryapi.com/products/json#sec-2.prs
type Sound struct {
	Filename string `json:"audio"`
}
