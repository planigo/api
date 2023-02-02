package entities

type Service struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ShopID     string  `json:"shop_id"`
	Duration    int     `json:"duration"`
}
