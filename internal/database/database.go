package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/yourusername/stock-portfolio-app/internal/config"
)

var DB *sql.DB

// Initialize the PostgreSQL connection
func InitDB() {
    var err error
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.AppConfig.DBHost,
        config.AppConfig.DBPort,
        config.AppConfig.DBUser,
        config.AppConfig.DBPassword,
        config.AppConfig.DBName,
    )

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Cannot reach database: ", err)
    }

    log.Println("Connected to PostgreSQL database")
}

// RunMigrations creates the tables if they don't exist
func RunMigrations() {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        phone VARCHAR(15) UNIQUE,
        password VARCHAR(255) NOT NULL
    );`

    _, err := DB.Exec(query)
    if err != nil {
        log.Fatalf("Error running migrations: %v", err)
    }

    log.Println("Migrations ran successfully")
}
