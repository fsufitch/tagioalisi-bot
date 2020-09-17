package types

type withCognateCrossReferences struct {
	CognateCrossReferences []CognateCrossReference `json:"cxs"`
}

// CognateCrossReference https://dictionaryapi.com/products/json#sec-2.cxs
type CognateCrossReference struct {
	Label   string                        `json:"cxl"`
	Targets []CognateCrossReferenceTarget `json:"cxtis"`
}

// CognateCrossReferenceTarget https://dictionaryapi.com/products/json#sec-2.cxs
type CognateCrossReferenceTarget struct {
	Label         string `json:"cxl"`
	TargetID      string `json:"cxr"`
	HyperlinkText string `json:"cxt"`
	SenseNumber   string `json:"cxn"`
}
