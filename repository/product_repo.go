package repository

import (
	"api/dao"
	"api/db"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// CreateProduct creates a product in the database
func CreateProduct(product *dao.Product) (*dao.Product, error) {
	product.ID = uuid.New()
	res := db.GetDatabase.Create(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

// GetProduct returns a product from the database
func GetProduct(id uuid.UUID) (*dao.Product, error) {
	var product dao.Product
	res := db.GetDatabase.First(&product, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("product with id %s not found", id))
	}
	return &product, nil
}

// GetProducts returns all products from the database
func GetProducts(args *dao.PaginationArguments) ([]*dao.Product, error) {
	var products []*dao.Product
	res := db.GetDatabase.Offset(args.Offset).Limit(args.Limit).Find(&products)
	if res.Error != nil {
		return nil, errors.New("no products found")
	}
	return products, nil
}

// UpdateProduct updates a product in the database
func UpdateProduct(product *dao.Product) (*dao.Product, error) {
	var productToUpdate dao.Product
	result := db.GetDatabase.Model(&productToUpdate).Where("id = ?", product.ID).Updates(product)
	if result.RowsAffected == 0 {
		return &productToUpdate, errors.New("product not updated")
	}
	return product, nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id string) error {
	var deletedProduct dao.Product
	result := db.GetDatabase.Where("id = ?", id).Delete(&deletedProduct)
	if result.RowsAffected == 0 {
		return errors.New("product does not exist")
	}
	return nil
}
