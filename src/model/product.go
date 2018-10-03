package models

// PRODUCT MODELS

import (
	"github.com/jinzhu/gorm"
)

// Product struct models database
type Product struct {
	gorm.Model
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
}

// Products struct payload mapper
type Products struct {
	ProductName         string
	ProductDescription  string
	ProductCategoryName string
}
