package model

type Request struct {
	UserID     int `json:"user_id"`
	MenuItemID int `json:"menu_item_id"`
}

type Restaurant struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CashBalance float32 `json:"cash_balance"`
}
