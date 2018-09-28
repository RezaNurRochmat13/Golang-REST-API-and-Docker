package src

import (
	"log"

	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/product/models"
	"github.com/jinzhu/gorm"
)

// Global variable
var db *gorm.DB

// ProductMigration does init migration when table in database doesn't exisr
func ProductMigration() {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Product)
	db.AutoMigrate(&models.ProductCategory{})
}
