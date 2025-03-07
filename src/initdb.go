package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type dbConn struct {
	DB *sqlx.DB
}

func initDb() (*dbConn, error) {
	log.Printf("Initializing postgres database\n")

	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDb := os.Getenv("PG_DB")
	pgSSL := os.Getenv("PG_SSL")

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDb, pgSSL)
	db, err := sqlx.Open("postgres", pgConnString)

	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	for i := 1; i <= 5; i++ {
		if err := db.Ping(); err != nil {
			if i == 5 {
				return nil, fmt.Errorf("error connecting to db: %w", err)
			} else {
				fmt.Printf("Error connecting to db at try %d, sleep 3 seconds\n", i)
				time.Sleep(3 * time.Second)
			}
		}
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
		return fmt.Errorf("error closing Postgresql: %w", err)
	}

	return nil
}
