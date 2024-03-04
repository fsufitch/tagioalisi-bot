package types

import "github.com/pkg/errors"

type withUsages struct {
	// https://dictionaryapi.com/products/json#sec-2.usages
	UsageParagraphs []UsageParagraphs `json:"usages"`
}

// UsageParagraphs https://dictionaryapi.com/products/json#sec-2.usages
type UsageParagraphs struct {
	Label string             `json:"pl"`
	Text  UsageParagraphText `json:"pt"`
}

// UsageParagraphText https://dictionaryapi.com/products/json#sec-2.usages
type UsageParagraphText arrayContainer

// UsageParagraphTextElementType is an enum type for the types of elements in the paragraph
type UsageParagraphTextElementType int

// Values for UsageParagraphTextElementType
const (
	UsageParagraphTextElementTypeUnknown UsageParagraphTextElementType = iota
	UsageParagraphTextElementTypeText
	UsageParagraphTextElementTypeVerbalIllustration
	UsageParagraphTextElementTypeSeeAlso
)

// UsageParagraphTextElementTypeFromString returns a UsageParagraphTextElementTYpe from its string ID
func UsageParagraphTextElementTypeFromString(id string) UsageParagraphTextElementType {
	switch id {
	case "text":
		return UsageParagraphTextElementTypeText
	case "vis":
		return UsageParagraphTextElementTypeVerbalIllustration
	case "uarefs":
		return UsageParagraphTextElementTypeSeeAlso
	default:
		return UsageParagraphTextElementTypeUnknown
	}
}

func (t UsageParagraphTextElementType) String() string {
	return []string{"", "text", "vis", "uarefs"}[t]
}

// Contents returns a copied slice of the contents in the UsageParagraphText
func (upt UsageParagraphText) Contents() ([]UsageParagraphTextElement, error) {
	elements := []UsageParagraphTextElement{}
	for _, el := range upt {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := UsageParagraphTextElementTypeFromString(key)
		switch typ {
		case UsageParagraphTextElementTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			elements = append(elements, UsageParagraphTextElement{Type: typ, Text: &out})
		case UsageParagraphTextElementTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			elements = append(elements, UsageParagraphTextElement{Type: typ, VerbalIllustration: &out})
		case UsageParagraphTextElementTypeSeeAlso:
			var out []UsageSeeAlso
			err = el.UnmarshalValue(&out)
			elements = append(elements, UsageParagraphTextElement{Type: typ, SeeAlso: out})
		default:
			err = errors.New("unknown element type in run-in")
		}
	}
	return elements, nil
}

// UsageParagraphTextElement is an element of the UsageParagraphText container
type UsageParagraphTextElement struct {
	Type               UsageParagraphTextElementType
	Text               *string
	VerbalIllustration *VerbalIllustration
	SeeAlso            []UsageSeeAlso
}

// UsageSeeAlso is "uaref" here https://dictionaryapi.com/products/json#sec-2.usages
type UsageSeeAlso struct {
	Reference string `json:"uaref"`
}
