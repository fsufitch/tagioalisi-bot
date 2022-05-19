package types

import "github.com/pkg/errors"

type withUndefinedRunonText struct {
	UndefinedRunonText UndefinedRunonText `json:"utxt"`
}

// UndefinedRunonText https://dictionaryapi.com/products/json#sec-2.uros
type UndefinedRunonText arrayContainer

// UndefinedRunonTextElementType is an enum type for the types of elements in UndefinedRunonText
type UndefinedRunonTextElementType int

// Values for UndefinedRunonTextElementType
const (
	UndefinedRunonTextElementTypeUnknown = iota
	UndefinedRunonTextElementTypeVerbalIllustrations
	UndefinedRunonTextElementTypeUsageNotes
)

// UndefinedRunonTextElementTypeFromString returns a UndefinedRunonTextElementType from its string ID
func UndefinedRunonTextElementTypeFromString(id string) UndefinedRunonTextElementType {
	switch id {
	case "vis":
		return UndefinedRunonTextElementTypeVerbalIllustrations
	case "uns":
		return UndefinedRunonTextElementTypeUsageNotes
	default:
		return UndefinedRunonTextElementTypeUnknown
	}
}

func (t UndefinedRunonTextElementType) String() string {
	return []string{"", "vis", "uns"}[t]
}

// Contents returns a copied slice of the contents in the UndefinedRunonText
func (uro UndefinedRunonText) Contents() ([]UndefinedRunonTextElement, error) {
	elements := []UndefinedRunonTextElement{}
	for _, el := range uro {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := UndefinedRunonTextElementTypeFromString(key)
		switch typ {
		case UndefinedRunonTextElementTypeVerbalIllustrations:
			var out []VerbalIllustration
			err = el.UnmarshalValue(&out)
			elements = append(elements, UndefinedRunonTextElement{Type: typ, withVerbalIllustrations: withVerbalIllustrations{out}})
		case UndefinedRunonTextElementTypeUsageNotes:
			var out []UsageNote
			err = el.UnmarshalValue(&out)
			elements = append(elements, UndefinedRunonTextElement{Type: typ, withUsageNotes: withUsageNotes{out}})
		default:
			err = errors.New("unknown element type in verbal illustration")
		}
	}
	return elements, nil
}

// UndefinedRunonTextElement is an element of the SI container
type UndefinedRunonTextElement struct {
	Type UndefinedRunonTextElementType
	withVerbalIllustrations
	withUsageNotes
}
