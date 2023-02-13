package types

type withDividedSense struct {
	DividedSense *DividedSense `json:"sdsense"`
}

// DividedSense https://dictionaryapi.com/products/json#sec-2.sdsense
type DividedSense struct {
	Divider string `json:"sd"`
	withEtymology
	withInflections
	withGeneralLabels
	withPronounciations
	withSenseSpecificGrammaticalLabel
	withSubjectStatusLabels
	withDefiningText
}
