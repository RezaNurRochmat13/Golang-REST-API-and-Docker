package main

import (
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/product"
	"github.com/Golang-REST-API-Gorilla-Mux-and-Docker/src/productCategory"
)

// BASE MAIN APPLICATION
func main() {
	product.ProductMainModule()
	productCategory.ProductCategoryMainModule()
}
