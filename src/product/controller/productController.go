package controller

// PRODUCT CONTROLLERS

import (
	"log"
	"net/http"

	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/product/models"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// GetAllProducts function does get all products
func GetAllProducts(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var products []models.Products
	var countProduct int

	db.Find(&products)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD REQUEST"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": products})
	}
}
