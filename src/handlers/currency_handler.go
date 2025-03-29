package handlers

import (
	"currency-api/services"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetCurrencyRates(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("EXCHANGE_RATE_API_KEY")
	if apiKey == "" {
		http.Error(w, `{"error": "API key is missing"}`, http.StatusUnauthorized)
		return
	}

	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/currency/"), "/")
	if len(pathParts) != 2 {
		http.Error(w, `{"error": "Invalid URL format. Use /currency/{base}/{target}"}`, http.StatusBadRequest)
		return
	}
	baseCurrency := strings.ToUpper(pathParts[0])
	targetCurrency := strings.ToUpper(pathParts[1])

	if len(baseCurrency) != 3 || len(targetCurrency) != 3 {
		http.Error(w, `{"error": "Invalid currency format. Use 3-letter currency codes."}`, http.StatusBadRequest)
		return
	}

	service := services.NewCurrencyService(apiKey)
	rates, err := service.FetchRates(baseCurrency)
	if err != nil {
		log.Printf("Failed to fetch currency rates: %v", err)
		http.Error(w, `{"error": "Failed to fetch currency rates"}`, http.StatusInternalServerError)
		return
	}

	conversionRate, exists := rates.Rates[targetCurrency]
	if !exists {
		http.Error(w, `{"error": "Target currency not found"}`, http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"base":   baseCurrency,
		"target": targetCurrency,
		"rate":   conversionRate,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
