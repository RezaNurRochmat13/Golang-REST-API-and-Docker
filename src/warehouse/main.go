package main

import (
	"github.com/gin-gonic/gin"
)

// MAIN APPLICATION

func main() {
	routers := gin.Default()

	v1 := routers.Group("/v1/api/")
	{
		v1.GET("product", getAllProducts)
	}
	routers.Run()
}
