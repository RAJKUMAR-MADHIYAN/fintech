package database

import (
	"database/sql"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// PostgreSQL connection
func DBConnect() *sql.DB {
	connStr := "user=postgres dbname=blasterbalu password=balaspba sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error Connecting to database : %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error Pinging database : %s", err)
	}

	fmt.Println("Connected to PostgreSQL")
	return db
}

// getDB returns database instance
func GetDB() *sql.DB {
	if db == nil {

		db = DBConnect()
	}

	return db
}
