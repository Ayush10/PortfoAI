package models

type Portfolio struct {
	ID     uint    `json:"id"`
	UserID uint    `json:"user_id"`
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
}
