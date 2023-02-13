package types

type withVerbalIllustrations struct {
	VerbalIllustrations []VerbalIllustration `json:"vis"`
}

// VerbalIllustration https://dictionaryapi.com/products/json#sec-2.vis
type VerbalIllustration struct {
	Text             string            `json:"t"`
	QuoteAttribution *QuoteAttribution `json:"aq"`
}

// ##### Below is a misguided arrayContainer-based implementation of vis

// // VerbalIllustration https://dictionaryapi.com/products/json#sec-2.vis
// type VerbalIllustration arrayContainer

// // VerbalIllustrationElementType is an enum type for the types of elements in VerbalIllustration
// type VerbalIllustrationElementType int

// // Values for VerbalIllustrationElementType
// const (
// 	VerbalIllustrationElementTypeUnknown = iota
// 	VerbalIllustrationElementTypeText
// 	VerbalIllustrationElementTypeAuthorQuotation
// )

// // VerbalIllustrationElementTypeFromString returns a VerbalIllustrationElementType from its string ID
// func VerbalIllustrationElementTypeFromString(id string) VerbalIllustrationElementType {
// 	switch id {
// 	case "t":
// 		return VerbalIllustrationElementTypeText
// 	case "aq":
// 		return VerbalIllustrationElementTypeAuthorQuotation
// 	default:
// 		return VerbalIllustrationElementTypeUnknown
// 	}
// }

// func (t VerbalIllustrationElementType) String() string {
// 	return []string{"", "t", "aq"}[t]
// }

// // Contents returns a copied slice of the contents in the VerbalIllustration
// func (vi VerbalIllustration) Contents() ([]VerbalIllustrationElement, error) {
// 	elements := []VerbalIllustrationElement{}
// 	for _, el := range vi {
// 		key, err := el.Key()
// 		if err != nil {
// 			return nil, err
// 		}
// 		typ := VerbalIllustrationElementTypeFromString(key)
// 		switch typ {
// 		case VerbalIllustrationElementTypeText:
// 			var out string
// 			err = el.UnmarshalValue(&out)
// 			elements = append(elements, VerbalIllustrationElement{Type: typ, Text: &out})
// 		case VerbalIllustrationElementTypeAuthorQuotation:
// 			var out QuoteAttribution
// 			err = el.UnmarshalValue(&out)
// 			elements = append(elements, VerbalIllustrationElement{Type: typ, QuoteAttribution: &out})
// 		default:
// 			err = errors.New("unknown element type in verbal illustration")
// 		}
// 	}
// 	return elements, nil
// }

// // VerbalIllustrationElement is an element of the SI container
// type VerbalIllustrationElement struct {
// 	Type             VerbalIllustrationElementType
// 	Text             *string
// 	QuoteAttribution *QuoteAttribution
// }
