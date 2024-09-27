package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/stock-portfolio-app/internal/services"
)

type LoginRequest struct {
    EmailOrPhone string `json:"emailOrPhone" binding:"required"`
    Password     string `json:"password" binding:"required"`
}

// Login an existing user
func Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := services.Authenticate(req.EmailOrPhone, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
