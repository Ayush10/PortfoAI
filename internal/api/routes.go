package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Ayush10/PortfoAI/internal/api/handlers"
)

func RegisterRoutes(router *gin.Engine) {
    v1 := router.Group("/api/v1")
    {
        v1.POST("/register", handlers.Register)  // Register handler
        v1.POST("/login", handlers.Login)        // Login handler
    }
}
