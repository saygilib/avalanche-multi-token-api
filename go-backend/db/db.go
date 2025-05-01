package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS api_logs (
		id SERIAL PRIMARY KEY,
		endpoint TEXT NOT NULL,
		request_payload TEXT NOT NULL,
		response TEXT NOT NULL,
		status INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,
	}

	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	}
	log.Println("✅ Tables created or already exist")
}
