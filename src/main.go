package main

import (
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/controller"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

func main() {
	BaseRoutes()

}

// BaseRoutes function does configuration your route apps
func BaseRoutes() {
	BaseRoutes := gin.Default()

	v1 := BaseRoutes.Group("/v1/api/")
	{
		v1.GET("product", controller.GetAllProducts)
		v1.GET("product/:ProductCode", controller.GetSingleProducts)
		v1.POST("product", controller.CreateProducts)
		v1.PUT("product/:ProductCode", controller.UpdateProducts)
		v1.DELETE("product/:ProductCode", controller.DeleteProducts)
	}

	BaseRoutes.Run(":8000")
}
