package types

import "github.com/pkg/errors"

type withEtymology struct {
	Etymology Etymology `json:"et"`
}

// Etymology https://dictionaryapi.com/products/json#sec-2.et
type Etymology arrayContainer

// EtymologyElementType is an enum type for the types of elements in Etymology
type EtymologyElementType int

// Values for EtymologyElementType
const (
	EtymologyElementTypeUnknown = iota
	EtymologyElementTypeText
	EtymologyElementTypeSupplementalInfo
)

// EtymologyElementTypeFromString returns a EtymologyElementType from its string ID
func EtymologyElementTypeFromString(id string) EtymologyElementType {
	switch id {
	case "text":
		return EtymologyElementTypeText
	case "et_snote":
		return EtymologyElementTypeSupplementalInfo
	default:
		return EtymologyElementTypeUnknown
	}
}

func (t EtymologyElementType) String() string {
	return []string{"", "text", "et_snote"}[t]
}

// Contents returns a copied slice of the contents in the Etymology
func (ety Etymology) Contents() ([]EtymologyElement, error) {
	elements := []EtymologyElement{}
	for _, el := range ety {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := EtymologyElementTypeFromString(key)
		switch typ {
		case EtymologyElementTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			elements = append(elements, EtymologyElement{Type: typ, Text: &out})
		case EtymologyElementTypeSupplementalInfo:
			var out SupplementalInfo
			err = el.UnmarshalValue(&out)
			elements = append(elements, EtymologyElement{Type: typ, SupplementalInfo: &out})
		default:
			err = errors.New("unknown element type in etymology")
		}
	}
	return elements, nil
}

// EtymologyElement is an element of the SI container
type EtymologyElement struct {
	Type             EtymologyElementType
	Text             *string
	SupplementalInfo *SupplementalInfo
}
