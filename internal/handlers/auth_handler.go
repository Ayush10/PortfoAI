// internal/api/handlers/auth_handler.go

package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Ayush10/PortfoAI/internal/services"
)

type LoginRequest struct {
    EmailOrPhone string `json:"emailOrPhone" binding:"required"`
    Password     string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var loginReq LoginRequest
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Authenticate user and generate JWT token
    token, err := services.Authenticate(loginReq.EmailOrPhone, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Return the JWT token as a response
    c.JSON(http.StatusOK, gin.H{"token": token})
}