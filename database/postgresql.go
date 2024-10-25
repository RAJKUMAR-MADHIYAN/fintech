package database

import (
	"database/sql"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// PostgreSQL connection
func DBConnect(cfg *config.AppConfig) *sql.DB {

	connStr := fmt.Sprintf("host=localhost port=5432 user=postgres password=balaspba dbname=blasterbalu sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

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
		cfg := config.LoadConfig()
		db = DBConnect(cfg)
	}

	return db
}
