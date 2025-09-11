package model

type Restaurant struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CashBalance float32 `json:"cash_balance"`
}
