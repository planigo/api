package models

type Service struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Duration    int     `json:"duration"`
	ShopID      string  `json:"shop_id"`
}
