package models

import (
	"github.com/jinzhu/gorm"
)

// ProductCategory struct model
type ProductCategory struct {
	gorm.Model
	ProductCategoryCode string `json:"productCategoryCode"`
	ProductCategoryName string `json:"productCategoryName"`
}
