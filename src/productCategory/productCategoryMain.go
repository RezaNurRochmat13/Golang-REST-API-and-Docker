package productCategory

import (
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/productCategory/controller"
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/productCategory/models"
	"github.com/gin-gonic/gin"
)

// ProductCategoryMainModule function does config for routing
func ProductCategoryMainModule() {
	models.ProductCategoryMigration()
	productCategoryRouters := gin.Default()

	v1 := productCategoryRouters.Group("/v1/api/")
	{
		v1.GET("productCategory", controller.GetAllProductCategory)
		v1.GET("productCategory/:ProductCategoryCode", controller.GetSingleProductCategory)
		v1.POST("productCategory", controller.CreateProductCategory)
		v1.PUT("productCategory/:ProductCategoryCode", controller.UpdateProductCategory)
		v1.DELETE("productCategory/:ProductCategoryCode", controller.DeleteProductCategory)
	}

	productCategoryRouters.Run()

}
