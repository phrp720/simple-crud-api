package dao

import "github.com/google/uuid"

// Product represents a product in the inventory
// @Description Product represents a product in the inventory
type Product struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" gorm:"default:'No description'"`
	SKU          string    `json:"sku" binding:"required"`
	Image        string    `json:"image" gorm:"default:'noimage.png'"`
	Price        float64   `json:"price" gorm:"default:0.0"`
	Stock        int       `json:"stock" gorm:"default:-1"`
	Availability bool      `json:"availability"`
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

type ProductsList struct {
	Products []Product `json:"products" binding:"required"`
}

type PlainProduct struct {
	Product UpdateProduct `json:"product" binding:"required"`
}

type PaginationArguments struct {
	Limit  int
	Offset int
}
