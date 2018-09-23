package main

import (
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/warehouse/controller"
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/warehouse/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MAIN APPLICATION

func main() {
	models.InitialMigration()
	routers := gin.Default()

	v1 := routers.Group("/v1/api/")
	{
		v1.GET("product", controller.GetAllProducts)
		v1.GET("product/:productCode", controller.GetSingleProducts)
	}
	routers.Run()
}
