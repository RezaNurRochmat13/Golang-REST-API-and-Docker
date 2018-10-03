package controller

// PRODUCT CONTROLLERS

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-rest-api/src/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetAllProducts function does get all products
func GetAllProducts(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var products []models.Products
	var countProduct int

	db.Select("products.product_name, products.product_description, products_categories.product_category_name").Joins("JOIN products_categories ON products_categories.product_category_code=products.product_category_code").Find(&products)
	db.Table("products").Count(&countProduct)

	c.JSON(http.StatusOK, gin.H{
		"count": countProduct,
		"data":  products,
		"total": countProduct})
}

// GetSingleProducts function does get single data products
func GetSingleProducts(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ProductCode := c.Param("ProductCode")

	var products models.Products

	db.Where("product_code = ?", ProductCode).Find(&products)

	if db.Where("product_code = ?", ProductCode).Find(&products).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "NOT FOUND"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": products})
	}
}

// CreateProducts function does create new products
func CreateProducts(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var createProduct models.Product

	c.BindJSON(&createProduct)
	db.Save(&createProduct)

	c.JSON(http.StatusOK, gin.H{"message": "Inserted successfully"})
}

// UpdateProducts function does update products
func UpdateProducts(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ProductCode := c.Param("ProductCode")

	var updateProducts models.Product
	db.Where("product_code = ?", ProductCode).Find(&updateProducts)

	if db.Where("product_code = ?", ProductCode).Find(&updateProducts).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "NOT FOUND RECORD"})
	} else {
		c.BindJSON(&updateProducts)
		db.Save(&updateProducts)
		c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
	}
}

// DeleteProducts function does delete products
func DeleteProducts(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ProductCode := c.Param("ProductCode")

	var deleteProduct models.Product
	db.Where("product_code = ?", ProductCode).Find(&deleteProduct)

	if db.Where("product_code = ?", ProductCode).Find(&deleteProduct).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "NOT FOUND RESOURCE"})
	} else {
		db.Unscoped().Delete(&deleteProduct)
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	}

}
