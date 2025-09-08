package main

import (
	"fmt"
	"log"
	"net/http"
	"task/internal/db"
	"task/pkg/api"
)

func main() {
	// connect & ping pg db
	err := db.Connect("postgres://postgres:@localhost:5432/test?sslmode=disable")

	if err != nil {
		panic("Database connection failed")
	}

	defer db.DB.Close()

	err = db.Ping()

	if err != nil {
		panic("Database ping failed")
	}

	// create httpMux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("GET /user-restaurants", api.GetUserRestaurantsHandler)
	mux.HandleFunc("POST /purchase", api.PurchaseMenuItemHandler)

	// start listening
	fmt.Println("Starting server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
