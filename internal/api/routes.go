package api

import (
	"github.com/Ayush10/PortfoAI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", handlers.Register) // Register handler
		v1.POST("/login", handlers.Login)       // Login handler
		v1.GET("/stocks", handlers.GetStocks)   // Route for fetching stock data
	}
}
