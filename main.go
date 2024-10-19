package main

import (
	"log"
	"os"

	"github.com/Ayush10/PortfoAI/configs"
	"github.com/Ayush10/PortfoAI/internal/api"
	"github.com/Ayush10/PortfoAI/internal/database"
	"github.com/Ayush10/PortfoAI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()     // Load configurations (DB, etc.)
	db := database.InitDB()  // Initialize the DB connection and get the instance
	database.DB = db         // Set the global DB variable
	database.RunMigrations() // Run migrations to create tables

	r := gin.Default()
	api.RegisterRoutes(r, db) // Register API routes, passing the db instance

	// chat feature
	apiKey := os.Getenv("CHATGPT_API_KEY") // Retrieve API key from environment variables
	if apiKey == "" {
		log.Fatal("CHATGPT_API_KEY environment variable not set")
	}
	chatHandler := handlers.NewChatHandler(apiKey)

	// Register Chat Routing
	r.POST("/chat", func(c *gin.Context) {
		chatHandler.HandleChat(c.Writer, c.Request)
	})

	// Provide static files (chat interface)
	r.Static("/", "./static")

	log.Println("Starting server on port 8080...")
	r.Run(":8080") // Listen and serve on port 8080
}
