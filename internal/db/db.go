package db

import (
	"database/sql"
	"task/internal/logger"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(connectionString string) error {
	log := logger.GetLogger()

	var err error
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("db: Connection established")
	return nil
}

func Ping() error {
	log := logger.GetLogger()
	err := DB.Ping()

	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("db: Ping success")
	return nil
}
