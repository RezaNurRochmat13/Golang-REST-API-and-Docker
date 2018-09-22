package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()
	routers.GET("/hello", Hello)
	routers.Run(":8080")
}

// Hello func does view hello
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello"})
}
