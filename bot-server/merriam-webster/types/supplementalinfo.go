package types

import "github.com/pkg/errors"

// SupplementalInfo https://dictionaryapi.com/products/json#sec-2.snote
type SupplementalInfo arrayContainer

// SupplementalInfoElementType is an enum type for the types of elements in SI
type SupplementalInfoElementType int

// Values for SupplementalInfoElementType
const (
	SupplementalInfoElementTypeUnknown = iota
	SupplementalInfoElementTypeText
	SupplementalInfoElementTypeRunIn
	SupplementalInfoElementTypeVerbalIllustration
)

// SupplementalInfoElementTypeFromString returns a SupplementalInfoElementType from its string ID
func SupplementalInfoElementTypeFromString(id string) SupplementalInfoElementType {
	switch id {
	case "t":
		return SupplementalInfoElementTypeText
	case "ri":
		return SupplementalInfoElementTypeRunIn
	case "vis":
		return SupplementalInfoElementTypeVerbalIllustration
	default:
		return SupplementalInfoElementTypeUnknown
	}
}

func (t SupplementalInfoElementType) String() string {
	return []string{"", "t", "ri", "vis"}[t]
}

// Contents returns a copied slice of the contents in the SupplementalInfo
func (si SupplementalInfo) Contents() ([]SupplementalInfoElement, error) {
	elements := []SupplementalInfoElement{}
	for _, el := range si {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := SupplementalInfoElementTypeFromString(key)
		switch typ {
		case SupplementalInfoElementTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			elements = append(elements, SupplementalInfoElement{Type: typ, Text: &out})
		case SupplementalInfoElementTypeRunIn:
			var out RunIn
			err = el.UnmarshalValue(&out)
			elements = append(elements, SupplementalInfoElement{Type: typ, RunIn: &out})
		case SupplementalInfoElementTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			elements = append(elements, SupplementalInfoElement{Type: typ, VerbalIllustration: &out})
		default:
			err = errors.New("unknown element type in supplemental info")
		}
	}
	return elements, nil
}

// SupplementalInfoElement is an element of the SI container
type SupplementalInfoElement struct {
	Type               SupplementalInfoElementType
	Text               *string
	RunIn              *RunIn
	VerbalIllustration *VerbalIllustration
}
