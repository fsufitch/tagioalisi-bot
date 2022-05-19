package types

type withArt struct {
	// https://dictionaryapi.com/products/json#sec-2.art
	Art Art `json:"art"`
}

// Art https://dictionaryapi.com/products/json#sec-2.art
type Art struct {
	ID      string `json:"artid"`
	Caption string `json:"capt"`
}
