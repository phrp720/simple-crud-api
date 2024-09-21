package router

import (
	"api/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/products", handler.GetProducts)                 //Get all products
	r.GET("/products/:id", handler.GetProduct)              //Get a product
	r.POST("/products/create", handler.PostProduct)         //Create a product
	r.POST("/products/bulk", handler.PostProducts)          //Create multiple products
	r.PUT("/products/update/:id", handler.PutProduct)       //Update a product
	r.DELETE("/products/delete/:id", handler.DeleteProduct) //Delete a product
	return r
}
