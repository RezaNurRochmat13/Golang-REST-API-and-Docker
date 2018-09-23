package controller

// PRODUCT CONTROLLERS

import (
	"log"
	"net/http"

	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/warehouse/models"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

// GetAllProducts function does get all products
func GetAllProducts(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var products []models.Product

	db.Find(&products)

	log.Println(products)

	c.JSON(http.StatusOK, gin.H{
		"count": 1,
		"data":  products,
		"total": 1})
}

// GetSingleProducts function does get single data products
func GetSingleProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "product"})
}
