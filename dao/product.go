package dao

import "github.com/google/uuid"

// Product represents a product in the inventory
// @Description Represents a product in the inventory with all its details
type Product struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SKU          string    `json:"sku"`
	Image        string    `json:"image" gorm:"default:'noimage.png'"`
	Price        float64   `json:"price" gorm:"default:0.0"`
	Stock        int       `json:"stock" gorm:"default:0"`
	Availability bool      `json:"availability"`
} // @name BaseProduct

// UpdateProduct represents a product update in the inventory
// @Description Represents the details of a product that are to be updated in the inventory
type UpdateProduct struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	SKU          string  `json:"sku"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	Stock        *int    `json:"stock"`
	Availability bool    `json:"availability"`
} // @name UpdateProduct

// CreatedProduct represents a product in the inventory
// @Description Represents the details of a new product to be created in the inventory
type CreatedProduct struct {
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description"`
	SKU          string  `json:"sku" binding:"required"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Availability bool    `json:"availability"`
} // @name CreatedProduct

// CreatedProductList represents a list of products in the inventory and used to Create product/s
// @Description Represents a list of new products to be created in the inventory
type CreatedProductList struct {
	Products []CreatedProduct `json:"products" binding:"required,dive"`
} // @name Products

// ChangedProduct represents a product in the inventory and used to Change the product
// @Description Represents the details of a product to be changed in the inventory
type ChangedProduct struct {
	Product UpdateProduct `json:"product" binding:"required"`
} // @name ChangedProduct

type PaginationArguments struct {
	Limit  int
	Offset int
}
