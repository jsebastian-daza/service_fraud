package models

// CurrencyResponse represents the response from a currency exchange API,
// including exchange rates, base currency, and any errors encountered.
type CurrencyResponse struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
	Error     CurrencyApiError
}

// HasError checks if the CurrencyResponse contains any error.
func (i *CurrencyResponse) HasError() bool {
	return i.Error.Code != 0 || i.Error.Message != ""
}
