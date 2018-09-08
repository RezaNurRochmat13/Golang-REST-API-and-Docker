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

func initialDatabaseConnection() {
	db, err := gorm.Open("mysql", "root:reza@tcp(127.0.0.1:3306)/store")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

}

// MAIN FUNCTION
func main() {
	initialMigration()
	routers := mux.NewRouter()
	routers.HandleFunc("/products", GetAllProducts).Methods("GET")
	routers.HandleFunc("/products", CreateNewProducts).Methods("POST")
	fmt.Println("Server serve in port 8080")
	log.Fatal(http.ListenAndServe(":8080", routers))
}

// GetAllProducts function does get all products
func GetAllProducts(w http.ResponseWriter, req *http.Request) {
	// Intial configuration migrations
	initialMigration()

	var products []Product
	db.Find(&products)
	fmt.Println("{}", products)

	json.NewEncoder(w).Encode(products)
}

// CreateNewProducts function does create new products
func CreateNewProducts(w http.ResponseWriter, r *http.Request) {
	initialDatabaseConnection()

	vars := mux.Vars(r)
	productCode := vars["productCode"]
	productName := vars["productName"]
	productPrice := vars["productPrice"]

	db.Create(&Product{ProductCode: productCode, ProductName: productName, ProductPrice: productPrice})
	fmt.Fprintf(w, "New user successfully created")

}
