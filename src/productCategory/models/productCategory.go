package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// ProductCategory struct model
type ProductCategory struct {
	gorm.Model
	ProductCategoryCode string `json:"productCategoryCode"`
	ProductCategoryName string `json:"productCategoryName"`
}

// ProductCategoryMigration does init migration when table in database doesn't exisr
func ProductCategoryMigration() {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&ProductCategory{})
}
