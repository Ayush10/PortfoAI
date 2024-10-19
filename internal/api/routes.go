package api

import (
	"github.com/Ayush10/PortfoAI/internal/handlers"
	"github.com/Ayush10/PortfoAI/internal/repository"
	"github.com/Ayush10/PortfoAI/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize repository
	portfolioRepo := repository.NewPortfolioRepository(db)

	// Initialize service
	portfolioService := services.NewPortfolioService(portfolioRepo)

	// Initialize handler
	portfolioHandler := handlers.NewPortfolioHandler(portfolioService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", handlers.Register) // Register handler
		v1.POST("/login", handlers.Login)       // Login handler
		v1.GET("/stocks", handlers.GetStocks)   // Route for fetching stock data

		// Portfolio routing
		v1.GET("/portfolio/:userID", portfolioHandler.GetUserPortfolio)
		v1.POST("/portfolio", portfolioHandler.AddToPortfolio)
	}
}
