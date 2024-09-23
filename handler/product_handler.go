package handler

import (
	"api/dao"
	"api/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// PostProducts creates multiple products
// @Summary      Create multiple products
// @Description  Create new products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        products  body      []dao.Product  true  "Products"
// @Success      201       {object}  []dao.Product
// @Router       /products/create [post]
func PostProducts(ctx *gin.Context) {
	var products []dao.Product
	err := ctx.Bind(&products)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var createdProducts []dao.Product
	for _, product := range products {
		res, err := repository.CreateProduct(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		createdProducts = append(createdProducts, *res)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"products": createdProducts,
	})
}

// GetProducts returns all products
// @Summary      Get all products
// @Description  Retrieve all products with pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"
// @Param        limit  query     int  false  "Page size"
// @Success      200    {object}  []dao.Product
// @Router       /products [get]
func GetProducts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &dao.PaginationArguments{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	res, err := repository.GetProducts(args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"products": res,
	})
}

// GetProduct returns a product
// @Summary      Get a product
// @Description  Retrieve a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  dao.Product
// @Router       /products/{id} [get]
func GetProduct(ctx *gin.Context) {
	id, err := StrToUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}
	res, err := repository.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"product": res,
	})
}

// PutProduct updates a product
// @Summary      Update a product
// @Description  Update an existing product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      string      true  "Product ID"
// @Param        product  body      dao.Product  true  "Product"
// @Success      200      {object}  dao.Product
// @Router       /products/update/{id} [put]
func PutProduct(ctx *gin.Context) {
	var updatedProduct dao.Product
	err := ctx.Bind(&updatedProduct)
	fmt.Print(updatedProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := StrToUUID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}
	dbProduct, err := repository.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if updatedProduct.Name != "" {
		dbProduct.Name = updatedProduct.Name
	}
	if updatedProduct.Description != "" {
		dbProduct.Description = updatedProduct.Description
	}
	if updatedProduct.SKU != "" {
		dbProduct.SKU = updatedProduct.SKU
	}
	if updatedProduct.Image != "noimage.png" {
		dbProduct.Image = updatedProduct.Image
	}
	dbProduct.Price = updatedProduct.Price
	if updatedProduct.Stock != -1 {
		dbProduct.Stock = updatedProduct.Stock
	}
	dbProduct.Availability = updatedProduct.Availability
	res, err := repository.UpdateProduct(dbProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"task": res,
	})
}

// DeleteProduct deletes a product
// @Summary      Delete a product
// @Description  Delete a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  "{"message": ""product deleted successfully"}"
// @Failure      404  "{"error": "product does not exist"}"
// @Router       /products/delete/{id} [delete]
func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := repository.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}

// DeleteProducts deletes multiple products
// @Summary      Delete multiple products
// @Description  Delete multiple products by IDs
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        ids  body      []string  true  "Product IDs"
// @Success      200  {object}  map[string]string
// @Router       /products/bulk/delete [delete]
func DeleteProducts(ctx *gin.Context) {
	var ids []string
	err := ctx.BindJSON(&ids)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	for _, id := range ids {
		err := repository.DeleteProduct(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Error deleting product with id %s: %s", id, err.Error()),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "products deleted successfully",
	})
}

func StrToUUID(id string) (uuid.UUID, error) {
	ConvertedId, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}
	return ConvertedId, nil
}
