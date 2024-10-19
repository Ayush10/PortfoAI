package services

import (
	"github.com/Ayush10/PortfoAI/internal/models"
	"github.com/Ayush10/PortfoAI/internal/repository"
)

type PortfolioService struct {
	repo *repository.PortfolioRepository
}

func NewPortfolioService(repo *repository.PortfolioRepository) *PortfolioService {
	return &PortfolioService{repo: repo}
}

func (s *PortfolioService) GetUserPortfolio(userID uint) ([]models.Portfolio, error) {
	return s.repo.GetUserPortfolio(userID)
}

func (s *PortfolioService) AddToPortfolio(userID uint, symbol string, amount float64) error {
	portfolio := &models.Portfolio{
		UserID: userID,
		Symbol: symbol,
		Amount: amount,
	}
	return s.repo.AddToPortfolio(portfolio)
}
