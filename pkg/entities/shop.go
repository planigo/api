package entities

type Shop struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id,omitempty"`
	CategoryID  string `json:"category_id,omitempty"`
}

type ShopRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
	CategoryID  string `json:"category_id"`
}
