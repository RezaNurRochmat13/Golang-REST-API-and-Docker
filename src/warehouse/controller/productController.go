package controller

// PRODUCT CONTROLLERS

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllProducts function does get all products
func GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count": 1,
		"data":  "all product",
		"total": 1})
}

// GetSingleProducts function does get single data products
func GetSingleProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "product"})
}
