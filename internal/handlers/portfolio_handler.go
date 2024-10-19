package handlers

import (
	"net/http"
	"strconv"

	"github.com/Ayush10/PortfoAI/internal/services"
	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct {
	portfolioService *services.PortfolioService
}

func NewPortfolioHandler(portfolioService *services.PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{
		portfolioService: portfolioService,
	}
}

func (h *PortfolioHandler) GetUserPortfolio(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 64)
	portfolio, err := h.portfolioService.GetUserPortfolio(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, portfolio)
}

func (h *PortfolioHandler) AddToPortfolio(c *gin.Context) {
	var input struct {
		UserID uint    `json:"user_id" binding:"required"`
		Symbol string  `json:"symbol" binding:"required"`
		Amount float64 `json:"amount" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.portfolioService.AddToPortfolio(input.UserID, input.Symbol, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added to portfolio"})
}
