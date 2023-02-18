package types

import (
	"github.com/pkg/errors"
)

// ErrInvalidSenseSequence is an error when the sense sequence is structured wrong
var ErrInvalidSenseSequence = errors.New("invalid sense sequence")

type withSenseSequence struct {
	SenseSequence SenseSequence `json:"sseq"`
}

// SenseSequence https://dictionaryapi.com/products/json#sec-2.vis
type SenseSequence arrayContainer

// SenseSequenceElementType is an enum type for the types of elements in SenseSequence
type SenseSequenceElementType int

// Values for SenseSequenceElementType
const (
	SenseSequenceElementTypeUnknown = iota
	SenseSequenceElementTypeSense
	SenseSequenceElementTypeAbbreviatedSense
	SenseSequenceElementTypeBindingSubstitute
	SenseSequenceElementTypeSubSequence
	SenseSequenceElementTypeParenthesizedSequence
)

// SenseSequenceElementTypeFromString returns a SenseSequenceElementType from its string ID
func SenseSequenceElementTypeFromString(id string) SenseSequenceElementType {
	switch id {
	case "sense":
		return SenseSequenceElementTypeSense
	case "sen":
		return SenseSequenceElementTypeAbbreviatedSense
	case "bs":
		return SenseSequenceElementTypeBindingSubstitute
	case "pseq":
		return SenseSequenceElementTypeParenthesizedSequence
	default:
		return SenseSequenceElementTypeUnknown
	}
}

func (t SenseSequenceElementType) String() string {
	return []string{"", "sense", "sen", "bs", "", "pseq"}[t]
}

// Contents returns a copied slice of the contents in the SenseSequence
func (vi SenseSequence) Contents() ([]SenseSequenceElement, error) {

	elements := []SenseSequenceElement{}
	for _, el := range vi {
		if len(el) == 0 {
			return nil, errors.New("zero length element")
		}
		if _, ok := el[0].([]interface{}); ok {
			out := SenseSequence{}
			if err := reUnmarshal(el, &out); err != nil {
				return nil, errors.Wrap(err, "error composing subsequence")

			}
			elements = append(elements, SenseSequenceElement{Type: SenseSequenceElementTypeSubSequence, SubSequence: &out})
			continue
		}

		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := SenseSequenceElementTypeFromString(key)
		switch typ {
		case SenseSequenceElementTypeSense:
			var out Sense
			err = el.UnmarshalValue(&out)
			elements = append(elements, SenseSequenceElement{Type: typ, Sense: &out})
		case SenseSequenceElementTypeAbbreviatedSense:
			var out AbbreviatedSense
			err = el.UnmarshalValue(&out)
			elements = append(elements, SenseSequenceElement{Type: typ, AbbreviatedSense: &out})
		case SenseSequenceElementTypeBindingSubstitute:
			var out BindingSubstitute
			err = el.UnmarshalValue(&out)
			elements = append(elements, SenseSequenceElement{Type: typ, BindingSubstitute: &out})
		case SenseSequenceElementTypeParenthesizedSequence:
			var out SenseSequence
			err = el.UnmarshalValue(&out)
			elements = append(elements, SenseSequenceElement{Type: typ, ParenthesizedSequence: &out})
		default:
			err = errors.New("unknown element type in verbal illustration")
		}
	}
	return elements, nil
}

// SenseSequenceElement is an element of the SI container
type SenseSequenceElement struct {
	Type                  SenseSequenceElementType
	Sense                 *Sense
	AbbreviatedSense      *AbbreviatedSense
	BindingSubstitute     *BindingSubstitute
	SubSequence           *SenseSequence
	ParenthesizedSequence *SenseSequence
}
