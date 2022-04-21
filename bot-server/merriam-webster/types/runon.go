package types

type withDefinedRunOns struct {
	// https://dictionaryapi.com/products/json#sec-2.dros
	DefinedRunOns []DefinedRunOn `json:"dros"`
}

// DefinedRunOn https://dictionaryapi.com/products/json#sec-2.dros
type DefinedRunOn struct {
	Phrase string `json:"drp"`
	withDefinitions
	withEtymology
	withGeneralLabels
	withPronounciations
	withParenthesizedSubjectStatusLabel
	withSubjectStatusLabels
	withVariants
}

type withUndefinedRunOns struct {
	// https://dictionaryapi.com/products/json#sec-2.uros
}

// UndefinedRunOn https://dictionaryapi.com/products/json#sec-2.uros
type UndefinedRunOn struct {
	Word string `json:"ure"`
	withFunctionalLabel
	Texts struct{} `json:"utxt"`
}
