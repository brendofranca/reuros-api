package rates

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type CurrencyService struct {
	APIKey string
	Cache  *cache.Cache
}

func NewCurrencyService(apiKey string) *CurrencyService {
	return &CurrencyService{
		APIKey: apiKey,
		Cache:  cache.New(10*time.Minute, 15*time.Minute),
	}
}

func (s *CurrencyService) FetchRates(baseCurrency string) (*CurrencyRate, error) {

	if cachedRates, found := s.Cache.Get(baseCurrency); found {
		log.Println("Cache hit for", baseCurrency)
		return cachedRates.(*CurrencyRate), nil
	}

	log.Println("Fetching rates", baseCurrency)
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

	var rates CurrencyRate
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, fmt.Errorf("JSON decoding error: %w", err)
	}

	if len(rates.Rates) == 0 {
		return nil, errors.New("error: received empty exchange rates data")
	}

	s.Cache.Set(baseCurrency, &rates, 10*time.Minute)

	return &rates, nil
}
