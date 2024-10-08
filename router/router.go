package router

import (
	"api/handler"
	"github.com/gin-gonic/gin"

	_ "api/docs" // This is important to import your generated docs
)
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"       // swagger embed files

func InitRouter() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/products", handler.GetProducts)                   // Get all products
		apiV1.GET("/products/:id", handler.GetProduct)                // Get a product
		apiV1.POST("/products/create", handler.PostProducts)          // Create product/s
		apiV1.PUT("/products/update/:id", handler.PutProduct)         // Update a product
		apiV1.DELETE("/products/delete/:id", handler.DeleteProduct)   // Delete a product
		apiV1.DELETE("/products/bulk/delete", handler.DeleteProducts) // Delete a list of products
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
