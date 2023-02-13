package types

import "github.com/pkg/errors"

type withSynonyms struct {
	// https://dictionaryapi.com/products/json#sec-2.syns
	SynonymParagraphs []SynonymParagraph `json:"syns"`
}

// SynonymParagraph https://dictionaryapi.com/products/json#sec-2.syns
type SynonymParagraph struct {
	Label       string               `json:"pl"`
	Text        SynonymParagraphText `json:"pt"`
	SeeAlsoRefs []string             `json:"sarefs"`
}

// SynonymParagraphText https://dictionaryapi.com/products/json#sec-2.syns
type SynonymParagraphText arrayContainer

// SynonymParagraphTextElementType is an enum type for the types of elements in the paragraph
type SynonymParagraphTextElementType int

// Values for SynonymParagraphTextElementType
const (
	SynonymParagraphTextElementTypeUnknown SynonymParagraphTextElementType = iota
	SynonymParagraphTextElementTypeText
	SynonymParagraphTextElementTypeVerbalIllustration
)

// SynonymParagraphTextElementTypeFromString returns a SynonymParagraphTextElementTYpe from its string ID
func SynonymParagraphTextElementTypeFromString(id string) SynonymParagraphTextElementType {
	switch id {
	case "text":
		return SynonymParagraphTextElementTypeText
	case "vis":
		return SynonymParagraphTextElementTypeVerbalIllustration
	default:
		return SynonymParagraphTextElementTypeUnknown
	}
}

func (t SynonymParagraphTextElementType) String() string {
	return []string{"", "text", "vis"}[t]
}

// Contents returns a copied slice of the contents in the SynonymParagraphText
func (upt SynonymParagraphText) Contents() ([]SynonymParagraphTextElement, error) {
	elements := []SynonymParagraphTextElement{}
	for _, el := range upt {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := SynonymParagraphTextElementTypeFromString(key)
		switch typ {
		case SynonymParagraphTextElementTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			elements = append(elements, SynonymParagraphTextElement{Type: typ, Text: &out})
		case SynonymParagraphTextElementTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			elements = append(elements, SynonymParagraphTextElement{Type: typ, VerbalIllustration: &out})
		default:
			err = errors.New("unknown element type in run-in")
		}
	}
	return elements, nil
}

// SynonymParagraphTextElement is an element of the SynonymParagraphText container
type SynonymParagraphTextElement struct {
	Type               SynonymParagraphTextElementType
	Text               *string
	VerbalIllustration *VerbalIllustration
}
