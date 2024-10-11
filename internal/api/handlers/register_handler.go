package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Ayush10/PortfoAI/internal/models"
	"github.com/Ayush10/PortfoAI/internal/services"
)

type RegisterRequest struct {
    EmailOrPhone string `json:"emailOrPhone" binding:"required"`
    Password     string `json:"password" binding:"required"`
}

// Register a new user
func Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    user := &models.User{
        Email: req.EmailOrPhone,   // Treating it as email for simplicity
        Password: req.Password,    // Raw password
    }

    err := services.CreateUser(user)
    if err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
