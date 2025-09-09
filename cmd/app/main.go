package main

import (
	"net/http"
	"os"
	"task/internal/handler"
	"task/internal/repository"
	"task/pkg/logger"

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

	// connect & ping pg db
	err = repository.Connect(os.Getenv("DBURL"))

	if err != nil {
		panic("Database connection failed")
	}

	defer repository.DB.Close()

	err = repository.Ping()

	if err != nil {
		panic("Database ping failed")
	}

	// create httpMux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("GET /user/{id}/restaurants", handler.GetUserRestaurantsHandler)
	mux.HandleFunc("POST /purchase", handler.PurchaseMenuItemHandler)

	// start listening
	logger.Log.Info("Starting server at localhost:" + os.Getenv("SERVER_PORT"))
	logger.Log.Error(http.ListenAndServe("localhost:"+os.Getenv("SERVER_PORT"), mux).Error())
}
