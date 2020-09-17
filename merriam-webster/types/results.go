package types

// CollegiateResult is a result from a search through the M-W Collegiate dictionary
type CollegiateResult struct {
	withMetadata
	withHomograph
	withHeadwordInfo

	withFunctionalLabel
	withGeneralLabels
	withSubjectStatusLabels
	withParenthesizedSubjectStatusLabel
	withSenseSpecificInflectionPluralLabel // might not belong here?
	withSenseSpecificGrammaticalLabel      // might not belong here?

	withInflections
	withCognateCrossReferences

	withDefinitions

	// misc
	withDefinedRunOns
	withUndefinedRunOns
	withDirectionalCrossReferences
	withUsages
	withSynonyms
	withQuotes
	withTable

	withEtymology
	withFirstKnownDate

	withShortDefinitions
}
