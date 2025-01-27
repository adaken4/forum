package db

import (
	"database/sql"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import SQLite3 driver
)

var DB *sql.DB // Global variable to hold the database connection

// Init initializes the database connection and creates the necessary tables.
func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	// defer DB.Close()

	createTables()
}

// createTables reads SQL statements from a file and executes them to set up the database schema.
func createTables() {
	sqlFile, err := os.Open("internal/db/schema.sql")
	if err != nil {
		log.Fatalf("Failed to open schema file: %v\n", err)
	}
	defer sqlFile.Close()

	sqlBytes, err := io.ReadAll(sqlFile)
	if err != nil {
		log.Fatalf("Failed to read schema file: %v\n", err)
	}

	sqlStatements := string(sqlBytes)

	// Execute the SQL statements. If an error occurs, log it and terminate the program.
	if _, err := DB.Exec(sqlStatements); err != nil {
		log.Fatalf("Failed to execute statements: %v\nQuery: %s\n", err, sqlStatements)
	}

	log.Println("All tables created successfully.")
}

func CleanupExpiredSessions() {
	query := `DELETE FROM sessions WHERE expires_at < ?`
	_, err := DB.Exec(query, time.Now())
	if err != nil {
		log.Printf("Error cleaning up sessions: %v", err)
	}
}
