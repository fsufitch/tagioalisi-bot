package types

type withTable struct {
	// https://dictionaryapi.com/products/json#sec-2.table
	Table Table `json:"table"`
}

// Table https://dictionaryapi.com/products/json#sec-2.table
type Table struct {
	ID          string `json:"tableid"`
	DisplayName string `json:"displayname"`
}
