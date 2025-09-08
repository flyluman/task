package main

import (
	"net/http"
	"os"
	"task/internal/db"
	"task/internal/logger"
	"task/pkg/api"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		panic("load .env failed")
	}

	// initialize logger
	logger.Init()
	log := logger.GetLogger()

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
	log.Info("Starting server at localhost:" + os.Getenv("SERVER_PORT"))
	log.Error(http.ListenAndServe("localhost:"+os.Getenv("SERVER_PORT"), mux).Error())
}
