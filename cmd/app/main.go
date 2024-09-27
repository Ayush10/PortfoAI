package main

import (
	"log"

	"github.com/Ayush10/PortfoAI/internal/api"
	"github.com/Ayush10/PortfoAI/internal/config"
	"github.com/Ayush10/PortfoAI/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()          // Load configurations (DB, etc.)
    database.InitDB()            // Initialize the DB connection
    database.RunMigrations()     // Run migrations to create tables

    r := gin.Default()
    api.RegisterRoutes(r)        // Register API routes

    log.Println("Starting server on port 8080...")
    r.Run(":8080")               // Listen and serve on port 8080
}
