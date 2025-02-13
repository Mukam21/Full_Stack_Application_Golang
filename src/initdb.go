package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dbConn struct {
	DB *sqlx.DB
}

func initDB() (*dbConn, error) {
	log.Print("Initializing postgres database\n")

	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")
	pgSSL := os.Getenv("PG_SSL")

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDB, pgSSL)
	//pgConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", pgUser, pgPassword, pgHost, pgPort, pgDB, pgSSL)

	log.Printf("Connecting to Postgresql\n")
	db, err := sqlx.Open("postgres", pgConnString)

	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", err) // %w
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err) // %w
	}

	if db == nil {
		log.Println("db is nil")
	}

	return &dbConn{
		DB: db,
	}, nil
}

func (d *dbConn) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %v", err) // %w
	}

	return nil
}
