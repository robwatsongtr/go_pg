package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Init() (*sql.DB, error) {
	host := "localhost"
	port := 5433
	user := "postgres"
	password := "54321"
	dbname := "guitar-clients"
	sslmode := "disable"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to database on port", port)

	return db, nil
}
