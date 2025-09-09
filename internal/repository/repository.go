package repository

import (
	"database/sql"
	"task/pkg/logger"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(connectionString string) error {
	var err error
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	logger.Log.Info("db: Connection established")
	return nil
}

func Ping() error {
	err := DB.Ping()

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	logger.Log.Info("db: Ping success")
	return nil
}
