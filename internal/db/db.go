package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(connectionString string) {
	var err error
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("db: Connection established")
	}
}

func Ping() {
	err := DB.Ping()

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("db: Ping success")
	}
}
