package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatResponse(t *testing.T) {
	ipRes := IpApiResponse{
		IP:          "1.1.1.1",
		CountryName: "Test Country",
		Location: Location{
			Languages: []Language{
				{Name: "Spanish", Code: "ES"},
				{Name: "English", Code: "EN"},
			},
		},
		Latitude:  34.0,
		Longitude: -58.0,
	}

	countryRes := CountryResponse{
		ArrayResponse: []CountryResponseElement{
			{
				Currencies: map[string]Currency{
					"USD": {Name: "United States Dollar", Symbol: "$"},
				},
				Cca2:      "TC",
				Timezones: []string{"UTC-3"},
			},
		},
	}

	currencyRes := CurrencyResponse{
		Rates: map[string]float64{"USD": 1.0},
	}

	response := Response{}
	response.FormatResponse(ipRes, countryRes, currencyRes)

	// La prueba no valida la salida, se puede hacer redirigiendo la salida estándar.
}

func TestGetCurrencyRates(t *testing.T) {
	rates := CurrencyResponse{
		Rates: map[string]float64{
			"USD": 1.0,
			"EUR": 0.85,
		},
	}

	result := getCurrencyRates("EUR", rates)
	expected := (1.0 * ((1 / 0.85) * 1.0)) // 1 EUR = 1 / 0.85 USD

	assert.Equal(t, expected, result)
}

func TestParseOffset(t *testing.T) {
	tests := []struct {
		offsetStr string
		expectErr bool
	}{
		{"UTC+3:00", false},
		{"UTC-2:30", false},
		{"UTC+invalid", true},
		{"invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.offsetStr, func(t *testing.T) {
			_, err := parseOffset(tt.offsetStr)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCountryResponse_HasError(t *testing.T) {
	tests := []struct {
		response CountryResponse
		hasError bool
	}{
		{CountryResponse{Error: CountryApiError{Code: 0, Message: ""}}, false},
		{CountryResponse{Error: CountryApiError{Code: 1, Message: "Some error"}}, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("HasError: %v", tt.hasError), func(t *testing.T) {
			assert.Equal(t, tt.hasError, tt.response.HasError())
		})
	}
}
