package types

// Sense https://dictionaryapi.com/products/json#sec-2.sense
type Sense struct {
	SenseNumber string `json:"sn"`
	withDefiningText
	withInflections
	withGeneralLabels
	withPronounciations
	withSenseSpecificGrammaticalLabel
	withDividedSense
	withSubjectStatusLabels
	withVariants
}

// AbbreviatedSense is a Sense without definition text
type AbbreviatedSense Sense

// BindingSubstitute is a special, broad kind of sense
type BindingSubstitute Sense
