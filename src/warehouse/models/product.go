package models

// PRODUCT MODELS

import "github.com/jinzhu/gorm"

var db *gorm.DB

// Product struct
type Product struct {
	gorm.Model
	ProductCode        string
	ProductName        string
	ProductDescription string
}

// initialMigration does init migration when table in database doesn't exisr
func initialMigration() {
	db.AutoMigrate(&Product{})
}
