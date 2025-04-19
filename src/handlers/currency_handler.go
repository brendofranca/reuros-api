package handlers

import (
	"log"
	"net/http"
	"reuros-api/services"
	"reuros-api/utils"
	"strings"
)

// GetCurrencyRates godoc
// @Summary Get currency conversion rate
// @Description Get the conversion rate between two currencies
// @Tags currency
// @Accept  json
// @Produce  json
// @Param base path string true "Base currency code"
// @Param target path string true "Target currency code"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /currency/{base}/{target} [get]
func GetCurrencyRates(w http.ResponseWriter, r *http.Request, service *services.CurrencyService) {
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/currency/"), "/")
	if len(pathParts) != 2 {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid URL format. Use /currency/{base}/{target}",
		})
		return
	}
	baseCurrency := strings.ToUpper(pathParts[0])
	targetCurrency := strings.ToUpper(pathParts[1])

	if len(baseCurrency) != 3 || len(targetCurrency) != 3 {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid currency format. Use 3-letter currency codes.",
		})
		return
	}

	rates, err := service.FetchRates(baseCurrency)
	if err != nil {
		log.Printf("Failed to fetch currency rates: %v", err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch currency rates",
		})
		return
	}

	conversionRate, exists := rates.Rates[targetCurrency]
	if !exists {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{
			"error": "Target currency not found",
		})
		return
	}

	response := map[string]interface{}{
		"base":            baseCurrency,
		"target":          targetCurrency,
		"rate":            conversionRate,
		"last_update_utc": rates.Date,
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}
