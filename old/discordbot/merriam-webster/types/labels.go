package types

type withFunctionalLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.fl
	Function string `json:"fl"`
}

type withGeneralLabels struct {
	// https://dictionaryapi.com/products/json#sec-2.lbs
	Labels []string `json:"lbs"`
}

type withSubjectStatusLabels struct {
	// https://dictionaryapi.com/products/json#sec-2.sls
	SubjectStatusLabels []string `json:"sls"`
}

type withParenthesizedSubjectStatusLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.psl
	ParenthesizedSubjectStatusLabel string `json:"psl"`
}

type withSenseSpecificInflectionPluralLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.spl
	SenseSpecificInflectionPluralLabel string `json:"spl"`
}

type withSenseSpecificGrammaticalLabel struct {
	// https://dictionaryapi.com/products/json#sec-2.sgram
	SenseSpecificGrammaticalLabel string `json:"sgram"`
}
