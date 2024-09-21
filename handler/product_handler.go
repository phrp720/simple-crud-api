package handler

import (
	"api/dao"
	"api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func GetProducts(ctx *gin.Context) {
	res, err := repository.GetProducts()
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
