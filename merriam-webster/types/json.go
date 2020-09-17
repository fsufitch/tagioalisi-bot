package types

import "encoding/json"

func reUnmarshal(value interface{}, output interface{}) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return json.Unmarshal(encoded, output)
}
