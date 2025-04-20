package rates

type CurrencyRate struct {
	Base  string             `json:"base_code"`
	Date  string             `json:"time_last_update_utc"`
	Rates map[string]float64 `json:"conversion_rates"`
}
