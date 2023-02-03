package entities

type Hour struct {
	Id     string `json:"id"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Day    int    `json:"day"`
	ShopID string `json:"shop_id"`
}
