package rates

import (
	"log"
	"net/http"
	"reuros-api/pkg"
	"strings"
)

// GetCurrencyRates godoc
// @Summary Get currency conversion rate
// @Description Get the conversion rate between two currencies
// @Tags currency
// @Accept  json
// @Produce  json
// @Param base query string true "Base currency code"
// @Param target query string true "Target currency code"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /currency-rates [get]
func GetCurrencyRates(w http.ResponseWriter, r *http.Request, service *CurrencyService) {
	baseCurrency := strings.ToUpper(r.URL.Query().Get("base"))
	targetCurrency := strings.ToUpper(r.URL.Query().Get("target"))

	if len(baseCurrency) != 3 || len(targetCurrency) != 3 {
		pkg.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid currency format. Use 3-letter currency codes.",
		})
		return
	}

	rates, err := service.FetchRates(baseCurrency)
	if err != nil {
		log.Printf("Failed to fetch currency rates: %v", err)
		pkg.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch currency rates",
		})
		return
	}

	conversionRate, exists := rates.Rates[targetCurrency]
	if !exists {
		pkg.WriteJSONResponse(w, http.StatusNotFound, map[string]string{
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

	pkg.WriteJSONResponse(w, http.StatusOK, response)
}
