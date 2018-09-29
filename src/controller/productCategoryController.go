package controller

import (
	"log"
	"net/http"

	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAllProductCategory function does get all product category
func GetAllProductCategory(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var productCategory []models.ProductsCategory
	var countProductCategory int

	db.Find(&productCategory)
	db.Table("product_category").Count(&countProductCategory)

	c.JSON(http.StatusOK, gin.H{
		"total": countProductCategory,
		"data":  productCategory,
		"count": countProductCategory,
	})
}

// GetSingleProductCategory function does get specific row from product category
func GetSingleProductCategory(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ProductCategoryCode := c.Param("ProductCategoryCode")

	var categoryProduct models.ProductsCategory

	db.Where("product_category_code = ?", ProductCategoryCode).Find(&categoryProduct)

	if db.Where("product_code = ?", ProductCategoryCode).Find(&categoryProduct).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found %ProductCategoryCode"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": categoryProduct})
	}
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
