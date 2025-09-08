package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(connectionString string) error {
	var err error
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	fmt.Println("db: Connection established")
	return nil
}

func Ping() {
	err := DB.Ping()

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("db: Ping success")
	}
}
