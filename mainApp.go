package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Product struct
type Product struct {
	gorm.Model
	ProductCode  string
	ProductName  string
	ProductPrice string
}

// Declare global variable for DB
var db *gorm.DB

// initialMigration function does config database connection and migration
func initialMigration() {
	var err error
	db, err = gorm.Open("mysql", "root:reza@tcp(127.0.0.1:3306)/store")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Auto migration schema
	db.AutoMigrate(&Product{})
}

// MAIN FUNCTION
func main() {
	initialMigration()
	routers := mux.NewRouter()
	routers.HandleFunc("/products", GetAllProducts).Methods("GET")
	routers.HandleFunc("/products/{id}", GetDetailProducts).Methods("GET")
	routers.HandleFunc("/products", CreateNewProducts).Methods("POST")
	fmt.Println("Server serve in port 8080")
	log.Fatal(http.ListenAndServe(":8080", routers))
}

// GetAllProducts function does get all products
func GetAllProducts(w http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open("mysql", "root:reza@tcp(127.0.0.1:3306)/store")

	if err != nil {
		panic("failed to connect database")
	}

	var products []Product
	db.Find(&products)

	json.NewEncoder(w).Encode(products)
}

// GetDetailProducts function does get single records
func GetDetailProducts(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:reza@tcp(127.0.0.1:3306)/store")

	if err != nil {
		panic("failed to connect database")
	}

	params := mux.Vars(r)
	var products Product

	db.First(&products, params["id"])
	json.NewEncoder(w).Encode(&products)
}

// CreateNewProducts function does create new products
func CreateNewProducts(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:reza@tcp(127.0.0.1:3306)/store")

	if err != nil {
		panic("failed to connect database")
	}

	var products Product
	json.NewDecoder(r.Body).Decode(&products)

	db.Create(&products)
}
