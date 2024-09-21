package dao

type Product struct {
	ID          string  `json:"id" gorm:"primarykey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Availabilty bool    `json:"availabilty"`
}
