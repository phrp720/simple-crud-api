package dao

// Product represents a product in the inventory
// @Description Product represents a product in the inventory
type Product struct {
	ID           string  `json:"id" gorm:"primarykey"`
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" gorm:"default:'No description'"`
	SKU          string  `json:"sku" binding:"required"`
	Image        string  `json:"image" gorm:"default:'noimage.png'"`
	Price        float64 `json:"price" gorm:"default:0.0"`
	Stock        int     `json:"stock" gorm:"default:-1"`
	Availability bool    `json:"availabilty"`
} // @name Product

// UpdateProduct represents a product update in the inventory
// @Description UpdateProduct represents a product update in the inventory
type UpdateProduct struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	SKU          string  `json:"sku"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Availability bool    `json:"availability"`
} // @name UpdateProduct

type PaginationArguments struct {
	Limit  int32
	Offset int32
}
