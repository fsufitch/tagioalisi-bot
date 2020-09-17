package types

import "github.com/pkg/errors"

// RunIn https://dictionaryapi.com/products/json#sec-2.ri
type RunIn arrayContainer

// RunInElementType is an enum type for the types of elements in RI
type RunInElementType int

// Values for RuninElementType
const (
	RunInElementTypeUnknown RunInElementType = iota
	RunInElementTypeText
	RunInElementTypeRunInWrap
)

// RunInElementTypeFromString returns a RunInElementTYpe from its string ID
func RunInElementTypeFromString(id string) RunInElementType {
	switch id {
	case "text":
		return RunInElementTypeText
	case "riw":
		return RunInElementTypeRunInWrap
	default:
		return RunInElementTypeUnknown
	}
}

func (t RunInElementType) String() string {
	return []string{"", "text", "riw"}[t]
}

// Contents returns a copied slice of the contents in the RunIn
func (ri RunIn) Contents() ([]RunInElement, error) {
	elements := []RunInElement{}
	for _, el := range ri {
		key, err := el.Key()
		if err != nil {
			return nil, err
		}
		typ := RunInElementTypeFromString(key)
		switch typ {
		case RunInElementTypeText:
			var out string
			err = el.UnmarshalValue(&out)
			elements = append(elements, RunInElement{Type: typ, Text: &out})
		case RunInElementTypeRunInWrap:
			var out RunInWrap
			err = el.UnmarshalValue(&out)
			elements = append(elements, RunInElement{Type: typ, RunInWrap: &out})
		default:
			err = errors.New("unknown element type in run-in")
		}
	}
	return elements, nil
}

// RunInElement is an element of the RunIn container
type RunInElement struct {
	Type      RunInElementType
	Text      *string
	RunInWrap *RunInWrap
}

// RunInWrap https://dictionaryapi.com/products/json#sec-2.ri
type RunInWrap struct {
	Word string
	withPronounciations
	withVariants
}
