// internal/api/handlers/auth_handler.go

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

func Login(c *gin.Context) {
    var loginReq LoginRequest
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Check credentials using a service function
    user, err := services.AuthenticateUser(loginReq.EmailOrPhone, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT token (or session token)
    token, err := services.GenerateJWT(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
        return
    }

    // Return the JWT token as a response
    c.JSON(http.StatusOK, gin.H{"token": token})
}
