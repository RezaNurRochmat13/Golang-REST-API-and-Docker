package main

import (
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/controller"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	BaseRoutes()

}

// BaseRoutes function does configuration your route apps
func BaseRoutes() {
	BaseRoutes := gin.Default()

	v1 := BaseRoutes.Group("/v1/api/")
	{
		// Product main routing
		v1.GET("product", controller.GetAllProducts)
		v1.GET("product/:ProductCode", controller.GetSingleProducts)
		v1.POST("product", controller.CreateProducts)
		v1.PUT("product/:ProductCode", controller.UpdateProducts)
		v1.DELETE("product/:ProductCode", controller.DeleteProducts)

		// Product category main routing
		v1.GET("product-category", controller.GetAllProductCategory)
		v1.GET("product-category/:ProductCategoryCode", controller.GetSingleProductCategory)

	}

	BaseRoutes.Run(":8000")
}
