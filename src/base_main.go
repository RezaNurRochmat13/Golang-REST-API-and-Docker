package main

import (
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/product/controller"
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/product/models"
	"github.com/gin-gonic/gin"
)

// BASE MAIN APPLICATION

func main() {

	// PRODUCT MODULE
	productModule()
}

// PRODUCT MODULE ROUTING
func productModule() {
	models.InitialMigration()
	routers := gin.Default()

	v1 := routers.Group("/v1/api/")
	{
		v1.GET("product", controller.GetAllProducts)
		v1.GET("product/:ProductCode", controller.GetSingleProducts)
		v1.POST("/product", controller.CreateProducts)
		v1.PUT("/product/:ProductCode", controller.UpdateProducts)
		v1.DELETE("/product/:ProductCode", controller.DeleteProducts)
	}

	routers.Run()
}
