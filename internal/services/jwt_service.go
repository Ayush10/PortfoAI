package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/Ayush10/PortfoAI/internal/models"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(user *models.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": user.ID,
        "exp":    time.Now().Add(time.Hour * 24).Unix(), // 1-day expiration
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
