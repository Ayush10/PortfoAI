package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ayush10/PortfoAI/internal/models"
)

const yahooFinanceAPIBaseURL = "https://query1.finance.yahoo.com/v8/finance/chart/"

type Stock struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
}

type YahooFinanceClient struct {
	httpClient *http.Client
}

func NewYahooFinanceClient() *YahooFinanceClient {
	return &YahooFinanceClient{
		httpClient: &http.Client{},
	}
}

func (c *YahooFinanceClient) FetchStockData(symbols []string) ([]models.Stock, error) {
	var stocks []models.Stock

	for _, symbol := range symbols {
		stock, err := c.fetchSingleStockData(symbol)
		if err != nil {
			fmt.Printf("Error fetching data for %s: %v\n", symbol, err)
			continue
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (c *YahooFinanceClient) fetchSingleStockData(symbol string) (models.Stock, error) {
	url := fmt.Sprintf("%s%s?interval=1d&range=1d", yahooFinanceAPIBaseURL, symbol)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return models.Stock{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Stock{}, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Stock{}, err
	}

	chart, ok := result["chart"].(map[string]interface{})
	if !ok {
		return models.Stock{}, fmt.Errorf("invalid response structure")
	}

	results, ok := chart["result"].([]interface{})
	if !ok || len(results) == 0 {
		return models.Stock{}, fmt.Errorf("no results found")
	}

	firstResult, ok := results[0].(map[string]interface{})
	if !ok {
		return models.Stock{}, fmt.Errorf("invalid result structure")
	}

	meta, ok := firstResult["meta"].(map[string]interface{})
	if !ok {
		return models.Stock{}, fmt.Errorf("meta data not found")
	}

	regularMarketPrice, ok := meta["regularMarketPrice"].(float64)
	if !ok {
		return models.Stock{}, fmt.Errorf("regular market price not found")
	}

	previousClose, ok := meta["previousClose"].(float64)
	if !ok {
		return models.Stock{}, fmt.Errorf("previous close not found")
	}

	change := regularMarketPrice - previousClose

	return models.Stock{
		Symbol: symbol,
		Price:  regularMarketPrice,
		Change: change,
	}, nil
}
