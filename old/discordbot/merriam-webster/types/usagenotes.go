package types

import "github.com/pkg/errors"

type withUsageNotes struct {
	UsageNotes []UsageNote
}

// UsageNote // https://dictionaryapi.com/products/json#sec-2.uns
type UsageNote arrayContainer

// UsageNoteElementType is an enum type for the types of elements in the Usage Note
type UsageNoteElementType int

// Values for UsageNoteElementType
const (
	UsageNoteElementTypeUnknown UsageNoteElementType = iota
	UsageNoteElementTypeText
	UsageNoteElementTypeRunIn
	UsageNoteElementTypeVerbalIllustration
)

// UsageNoteElementTypeFromString returns a UsageNoteElementTYpe from its string ID
func UsageNoteElementTypeFromString(id string) UsageNoteElementType {
	switch id {
	case "text":
		return UsageNoteElementTypeText
	case "ri":
		return UsageNoteElementTypeRunIn
	case "vis":
		return UsageNoteElementTypeVerbalIllustration
	default:
		return UsageNoteElementTypeUnknown
	}
}

func (t UsageNoteElementType) String() string {
	return []string{"", "text", "ri", "vis"}[t]
}

// Contents returns a copied slice of the contents in the UsageNote
func (un UsageNote) Contents() ([]UsageNoteElement, error) {
	elements := []UsageNoteElement{}
	for _, el := range un {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := UsageNoteElementTypeFromString(key)
		switch typ {
		case SupplementalInfoElementTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			elements = append(elements, UsageNoteElement{Type: typ, Text: &out})
		case SupplementalInfoElementTypeRunIn:
			var out RunIn
			err = el.UnmarshalValue(&out)
			elements = append(elements, UsageNoteElement{Type: typ, RunIn: &out})
		case SupplementalInfoElementTypeVerbalIllustration:
			var out VerbalIllustration
			err = el.UnmarshalValue(&out)
			elements = append(elements, UsageNoteElement{Type: typ, VerbalIllustration: &out})
		default:
			err = errors.New("unknown element type in supplemental info")
		}
	}
	return elements, nil
}

// UsageNoteElement is an element of the UsageNote container
type UsageNoteElement struct {
	Type               UsageNoteElementType
	Text               *string
	RunIn              *RunIn
	VerbalIllustration *VerbalIllustration
}
