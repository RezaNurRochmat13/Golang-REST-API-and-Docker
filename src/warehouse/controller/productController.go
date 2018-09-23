package controller

// PRODUCT CONTROLLERS

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count": 1,
		"data":  "all product",
		"total": 1})
}
