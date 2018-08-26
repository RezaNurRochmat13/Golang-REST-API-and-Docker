package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routers := mux.NewRouter()
	routers.HandleFunc("/Hello", Hello)
	fmt.Println("Server serve in port 8080")
	log.Fatal(http.ListenAndServe(":8080", routers))
}

// Hello function
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello golang")
}
