package database

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connString := "user=postgres password=Norman@123.9 dbname=nof_v1_db sslmode=disable"
	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error opening database %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error Connecting to the database %v", err)
	}

	fmt.Println("Successfully Connected!")
}
