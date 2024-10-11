package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize the SQLite connection
func InitDB() {
    var err error
    
    // Ensure that the data directory exists
    dbDir := "./data"
    if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
        log.Fatal("Error creating database directory: ", err)
    }

    dbPath := filepath.Join(dbDir, "portfoai.db")
    DB, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatal("Error opening database: ", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Cannot connect to database: ", err)
    }

    log.Println("Connected to SQLite database")
}

// RunMigrations creates the tables if they don't exist
func RunMigrations() {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT UNIQUE NOT NULL,
            phone TEXT UNIQUE,
            password TEXT NOT NULL
        );`,
        `CREATE TABLE IF NOT EXISTS portfolios (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            asset_name TEXT NOT NULL,
            asset_type TEXT NOT NULL,
            quantity REAL NOT NULL,
            purchase_price REAL NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );`,
    }

    for _, query := range queries {
        _, err := DB.Exec(query)
        if err != nil {
            log.Fatalf("Error running migration: %v", err)
        }
    }

    log.Println("Migrations ran successfully")
}