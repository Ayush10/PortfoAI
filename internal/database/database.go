package database

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize the SQLite connection
func InitDB() *gorm.DB {
	var err error

	// Ensure that the data directory exists
	dbDir := "./data"
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatal("Error creating database directory: ", err)
	}

	dbPath := filepath.Join(dbDir, "portfoai.db")
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	log.Println("Connected to SQLite database")
	return DB
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
		err := DB.Exec(query).Error
		if err != nil {
			log.Fatalf("Error running migration: %v", err)
		}
	}

	log.Println("Migrations ran successfully")
}
