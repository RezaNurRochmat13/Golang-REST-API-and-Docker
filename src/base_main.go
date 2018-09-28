package main

import (
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/controller"
	"github.com/gin-gonic/gin"
)

// BASE MAIN APPLICATION
func main() {
	baseRoutes()
}

func baseRoutes() {
	baseRoutes := gin.Default()

	v1 := baseRoutes.Group("/v1/api/")
	{
		v1.GET("product", controller.GetAllProducts)

		v1.GET("/productCategory", controler.GetSingleProductCategory)
	}
}
