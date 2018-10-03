package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-rest-api/src/model"
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

	if db.Where("product_category_code = ?", ProductCategoryCode).Find(&categoryProduct).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found %ProductCategoryCode"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": categoryProduct})
	}
}

// CreateProductCategory function does create new product category
func CreateProductCategory(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var createCategoryProduct models.ProductsCategory

	c.BindJSON(&createCategoryProduct)
	db.Save(&createCategoryProduct)
	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}

// UpdateProductCategory function does update product category
func UpdateProductCategory(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ProductCategoryCode := c.Param("ProductCategoryCode")

	var updateProductCategory models.ProductCategory

	db.Where("ProductCategoryCode = ?", ProductCategoryCode).Find(&updateProductCategory)

	if db.Where("product_category_code = ?", ProductCategoryCode).Find(&updateProductCategory).RecordNotFound() {
		c.JSON(http.StatusOK, gin.H{"error": "Resource not found"})
	} else {
		c.BindJSON(&updateProductCategory)
		c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
	}
}

// DeleteProductCategory function does delete product category
func DeleteProductCategory(c *gin.Context) {

	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ProductCategoryCode := c.Param("ProductCategoryCode")

	var deleteProductCategory models.ProductCategory

	db.Where("product_category_code = ?", ProductCategoryCode).Find(&deleteProductCategory)

	if db.Where("product_category_code = ?", ProductCategoryCode).Find(&deleteProductCategory).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
	} else {
		db.Unscoped().Delete(&deleteProductCategory)
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	}
}
