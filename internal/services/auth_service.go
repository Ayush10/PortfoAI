package services

import (
	"errors"

	"github.com/Ayush10/PortfoAI/internal/models"
	"github.com/Ayush10/PortfoAI/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CreateUser(user *models.User) error {
    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
    return repository.InsertUser(user)
}

func Authenticate(emailOrPhone, password string) (string, error) {
    user, err := repository.GetUserByEmailOrPhone(emailOrPhone)
    if err != nil {
        return "", errors.New("user not found")
    }

    if !CheckPasswordHash(password, user.Password) {
        return "", errors.New("incorrect password")
    }

    return GenerateJWT(user)
}
