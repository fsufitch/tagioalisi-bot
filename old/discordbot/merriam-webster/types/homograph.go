package types

type withHomograph struct {
	// https://dictionaryapi.com/products/json#sec-2.hom
	Homograph int `json:"hom"`
}
