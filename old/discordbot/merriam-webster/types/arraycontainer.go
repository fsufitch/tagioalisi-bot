package types

import (
	"github.com/pkg/errors"
)

// ErrInvalidArrayContainer is an error used when the array container is improperly structured
var ErrInvalidArrayContainer = errors.New("invalid array container")

// ArrayContainer is an abstract type for de/serializing a JSON structure shaped like:
// [ [string, any] ]
type arrayContainer []arrayContainerElement

func (a arrayContainer) Filter(key string) []*arrayContainerElement {
	result := []*arrayContainerElement{}
	for _, el := range a {
		if k, _ := el.Key(); k == key {
			result = append(result, &el)
		}
	}
	return result
}

type arrayContainerElement []interface{}

func (el arrayContainerElement) Key() (string, error) {
	if len(el) != 2 {
		return "", errors.Wrap(ErrInvalidArrayContainer, "container element must have length 2")
	}
	if key, ok := el[0].(string); ok {
		return key, nil
	}
	return "", errors.Wrap(ErrInvalidArrayContainer, "element[0] is not a string")
}

func (el arrayContainerElement) Value() (interface{}, error) {
	if len(el) != 2 {
		return "", errors.Wrap(ErrInvalidArrayContainer, "container element must have length 2")
	}
	return el[1], nil
}

func (el arrayContainerElement) UnmarshalValue(output interface{}) error {
	value, err := el.Value()
	if err != nil {
		return err
	}

	return reUnmarshal(value, output)
}
