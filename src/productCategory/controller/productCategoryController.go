package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProductCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": 1,
		"data":  "All Category Products",
		"count": 1,
	})
}

func GetSingleProductCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Single Categoryproducts"})
}

func CreateProductCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}

func UpdateProductCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func DeleteProductCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"messsage": "Deleted successfully"})
}
