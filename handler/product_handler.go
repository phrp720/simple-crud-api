package handler

import (
	"api/dao"
	"api/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PostProduct creates a product (This method is not used in the project)
func PostProduct(ctx *gin.Context) {
	var product dao.Product
	err := ctx.Bind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := repository.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"product": res,
	})
}

// PostProducts creates multiple products
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
func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
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
func PutProduct(ctx *gin.Context) {
	var updatedProduct dao.Product
	err := ctx.Bind(&updatedProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbProduct, err := repository.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbProduct.Name = updatedProduct.Name
	dbProduct.Description = updatedProduct.Description

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
