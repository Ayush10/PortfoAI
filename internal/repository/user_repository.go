package repository

import (
	"database/sql"
	"errors"

	"github.com/yourusername/stock-portfolio-app/internal/database"
	"github.com/yourusername/stock-portfolio-app/internal/models"
)

func GetUserByEmailOrPhone(emailOrPhone string) (*models.User, error) {
    user := &models.User{}

    query := `SELECT id, email, password FROM users WHERE email=$1 OR phone=$2`
    err := database.DB.QueryRow(query, emailOrPhone, emailOrPhone).Scan(&user.ID, &user.Email, &user.Password)
    if err == sql.ErrNoRows {
        return nil, errors.New("user not found")
    }
    return user, err
}

func InsertUser(user *models.User) error {
    query := `INSERT INTO users (email, phone, password) VALUES ($1, $2, $3)`
    _, err := database.DB.Exec(query, user.Email, user.Phone, user.Password)
    return err
}
