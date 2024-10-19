package repository

import (
	"github.com/Ayush10/PortfoAI/internal/models"
	"gorm.io/gorm"
)

type PortfolioRepository struct {
	DB *gorm.DB
}

func NewPortfolioRepository(db *gorm.DB) *PortfolioRepository {
	return &PortfolioRepository{DB: db}
}

func (r *PortfolioRepository) GetUserPortfolio(userID uint) ([]models.Portfolio, error) {
	var portfolio []models.Portfolio
	result := r.DB.Where("user_id = ?", userID).Find(&portfolio)
	return portfolio, result.Error
}

func (r *PortfolioRepository) AddToPortfolio(portfolio *models.Portfolio) error {
	return r.DB.Create(portfolio).Error
}
