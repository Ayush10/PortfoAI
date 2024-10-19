package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type YahooFinanceClient struct {
}

func NewYahooFinanceClient() *YahooFinanceClient {
	return &YahooFinanceClient{}
}

func (c *YahooFinanceClient) FetchStockData(symbols []string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{"symbol": "AAPL", "price": 150},
		{"symbol": "GOOGL", "price": 2800},
	}, nil
}

var yahooClient *YahooFinanceClient

func init() {
	yahooClient = NewYahooFinanceClient()
}

func GetStocks(c *gin.Context) {
	symbols := c.QueryArray("symbols")
	if len(symbols) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No stock symbols provided"})
		return
	}

	stocks, err := yahooClient.FetchStockData(symbols)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stock data"})
		return
	}
	c.JSON(http.StatusOK, stocks)
}
