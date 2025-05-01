package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type APILogEntry struct {
	Endpoint       string `json:"endpoint"`
	RequestPayload string `json:"requestPayload"`
	Response       string `json:"response"`
	Status         int    `json:"status"`
	Timestamp      string `json:"timestamp"`
}

func InitDB() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}
	log.Println("Connected to PostgreSQL")
}

func InsertAPILog(endpoint, requestPayload, response string, status int) {
	query := `
		INSERT INTO api_logs (endpoint, request_payload, response, status)
		VALUES ($1, $2, $3, $4)
	`

	_, err := DB.Exec(query, endpoint, requestPayload, response, status)
	if err != nil {
		log.Printf("Failed to insert API log: %v\n", err)
	}
}

func GetAllLogs() []APILogEntry {
	rows, err := DB.Query("SELECT endpoint, request_payload, response, status, created_at FROM api_logs ORDER BY id DESC")
	if err != nil {
		log.Printf("Failed to fetch API logs: %v\n", err)
		return nil
	}
	defer rows.Close()

	var logs []APILogEntry
	for rows.Next() {
		var l APILogEntry
		if err := rows.Scan(&l.Endpoint, &l.RequestPayload, &l.Response, &l.Status, &l.Timestamp); err != nil {
			log.Printf("Error scanning API log row: %v\n", err)
			continue
		}
		logs = append(logs, l)
	}
	return logs
}
