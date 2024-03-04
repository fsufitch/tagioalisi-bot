package types

import "github.com/pkg/errors"

type withDefiningText struct {
	// https://dictionaryapi.com/products/json#sec-2.dt
	DefiningText DefiningText `json:"dt"`
}

// DefiningTextElementType is an enum type for the types of elements in DT
type DefiningTextElementType int

// Values for DefiningTextElementType
const (
	DefiningTextElementTypeUnknown DefiningTextElementType = iota
	DefiningTextElementTypeText
	DefiningTextElementTypeBiography
	DefiningTextElementTypeCalledAlso
	DefiningTextElementTypeRunIn
	DefiningTextElementTypeSupplementalInfo
	DefiningTextElementTypeUsageNotes
	DefiningTextElementTypeVerbalIllustrations
)

// DefiningTextElementTypeFromString returns a DefiningTextElementType from its string ID
func DefiningTextElementTypeFromString(id string) DefiningTextElementType {
	switch id {
	case "text":
		return DefiningTextElementTypeText
	case "bnw":
		return DefiningTextElementTypeBiography
	case "ca":
		return DefiningTextElementTypeCalledAlso
	case "ri":
		return DefiningTextElementTypeRunIn
	case "snote":
		return DefiningTextElementTypeSupplementalInfo
	case "uns":
		return DefiningTextElementTypeUsageNotes
	case "vis":
		return DefiningTextElementTypeVerbalIllustrations
	default:
		return DefiningTextElementTypeUnknown
	}
}

func (t DefiningTextElementType) String() string {
	return []string{"", "text", "bnw", "ca", "ri", "snote", "uns", "vis"}[t]
}

// DefiningText https://dictionaryapi.com/products/json#sec-2.dt
type DefiningText arrayContainer

// Contents returns a copied slice of the contents in the DefiningText
func (dt DefiningText) Contents() ([]DefiningTextElement, error) {
	elements := []DefiningTextElement{}
	for _, el := range dt {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := DefiningTextElementTypeFromString(key)
		switch typ {
		case DefiningTextElementTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			elements = append(elements, DefiningTextElement{Type: typ, Text: &out})
		case DefiningTextElementTypeBiography:
			var out Biography
			err = el.UnmarshalValue(&out)
			elements = append(elements, DefiningTextElement{Type: typ, Biography: &out})
		case DefiningTextElementTypeCalledAlso:
			var out CalledAlso
			err = el.UnmarshalValue(&out)
			elements = append(elements, DefiningTextElement{Type: typ, CalledAlso: &out})
		case DefiningTextElementTypeRunIn:
			var out RunIn
			err = el.UnmarshalValue(&out)
			elements = append(elements, DefiningTextElement{Type: typ, RunIn: &out})
		case DefiningTextElementTypeSupplementalInfo:
			var out SupplementalInfo
			err = el.UnmarshalValue(&out)
			elements = append(elements, DefiningTextElement{Type: typ, SupplementalInfo: &out})
		case DefiningTextElementTypeUsageNotes:
			var out []UsageNote
			err = el.UnmarshalValue(&out)
			elements = append(elements, DefiningTextElement{Type: typ, withUsageNotes: withUsageNotes{out}})
		case DefiningTextElementTypeVerbalIllustrations:
			var out []VerbalIllustration
			err = el.UnmarshalValue(&out)
			elements = append(elements, DefiningTextElement{Type: typ, withVerbalIllustrations: withVerbalIllustrations{out}})
		default:
			err = errors.New("unknown element type in defining text")
		}
		if err != nil {
			return nil, err
		}
	}
	return elements, nil
}

// DefiningTextElement is an element in the DefiningText container.
// Type indicated which property is populated with data.
type DefiningTextElement struct {
	Type             DefiningTextElementType
	Text             *string
	Biography        *Biography
	CalledAlso       *CalledAlso
	RunIn            *RunIn
	SupplementalInfo *SupplementalInfo
	withUsageNotes
	withVerbalIllustrations
}

// Biography https://dictionaryapi.com/products/json#sec-2.bnw
type Biography struct {
	PersonalName  string `json:"pname"`
	Surname       string `json:"sname"`
	AlternateName string `json:"altname"`
}

// CalledAlso https://dictionaryapi.com/products/json#sec-2.ca
type CalledAlso struct {
	Intro   string             `json:"intro"`
	Targets []CalledAlsoTarget `json:"cats"`
}

// CalledAlsoTarget https://dictionaryapi.com/products/json#sec-2.ca
type CalledAlsoTarget struct {
	Text                string `json:"cat"`
	TargetID            string `json:"catref"`
	ParenthesizedNumber string `json:"pn"`
	withPronounciations
	withParenthesizedSubjectStatusLabel
}
