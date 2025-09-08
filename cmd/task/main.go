package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"task/internal/db"
	"task/pkg/api"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		panic("load .env failed")
	}

	// connect & ping pg db
	err = db.Connect(os.Getenv("DBURL"))

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
	mux.HandleFunc("GET /user/{id}/restaurants", api.GetUserRestaurantsHandler)
	mux.HandleFunc("POST /purchase", api.PurchaseMenuItemHandler)

	// start listening
	fmt.Println("Starting server at localhost:" + os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVER_PORT"), mux))
}
