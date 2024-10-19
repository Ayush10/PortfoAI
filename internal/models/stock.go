package models

type Stock struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
}
