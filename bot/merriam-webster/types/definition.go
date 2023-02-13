package types

type withDefinitions struct {
	Definitions []Definition `json:"def"`
}

// Definition https://dictionaryapi.com/products/json#sec-2.def
type Definition struct {
	VerbDivider string `json:"vd"`
	withSenseSequence
}
