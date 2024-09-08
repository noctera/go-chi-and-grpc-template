package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// DB is the global database connection pool
var DB *pgxpool.Pool

// Init initializes the database connection pool
func Init(databaseURL string) {
	var err error
	DB, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Optionally, ping the database to ensure the connection is established
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Connected to the database successfully")
}

// Close closes the database connection pool
func Close() {
	DB.Close()
}
