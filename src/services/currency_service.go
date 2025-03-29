package services

import (
	"currency-api/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CurrencyService struct {
	APIKey string
}

func NewCurrencyService(apiKey string) *CurrencyService {
	return &CurrencyService{APIKey: apiKey}
}

func (s *CurrencyService) FetchRates(baseCurrency string) (*models.CurrencyRate, error) {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", s.APIKey, baseCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch currency rates: HTTP %d, response: %s", resp.StatusCode, string(body))
	}

	var rates models.CurrencyRate
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, fmt.Errorf("JSON decoding error: %w", err)
	}

	if rates.Rates == nil || len(rates.Rates) == 0 {
		return nil, errors.New("error: received empty exchange rates data")
	}

	return &rates, nil
}
