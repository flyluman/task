package main

import (
	"database/sql"
	"net/http"
	"os"
	"task/internal/handler"
	"task/internal/repository"
	"task/internal/service"
	"task/pkg/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		panic("load .env failed")
	}

	// initialize logger
	logger.Init()

	// connect db
	db, err := sql.Open("postgres", os.Getenv("DBURL"))

	if err != nil {
		panic("Error connecting database")
	}

	defer db.Close()

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handler.UserHandler{UserService: service}

	// create httpMux
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("GET /user/{id}/restaurants", handler.GetUserRestaurantsHandler)
	mux.HandleFunc("POST /purchase", handler.PurchaseMenuItemHandler)

	// start listening
	logger.Log.Info("Starting server at localhost:" + os.Getenv("SERVER_PORT"))
	logger.Log.Error(http.ListenAndServe("localhost:"+os.Getenv("SERVER_PORT"), mux).Error())
}
