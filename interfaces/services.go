package interfaces

import "service_fraud/models"

type IpInformation interface {
	// Geolocation returns the geolocation data as an IpApiResponse for the specified IP.
	Geolocation(ip string) models.IpApiResponse
}

type CountryInformation interface {
	// GetCountryInformation returns the country information for the given country name.
	GetCountryInformation(country string) models.CountryResponse
}

type CurrencyInformation interface {
	// GetCurrencyInformation returns the currency information as a CurrencyResponse.
	GetCurrencyInformation() models.CurrencyResponse
}

type StatsInformation interface {
	// GetStats retrieves statistical data as a string.
	GetStats() string
	// Combine processes a StatsRequest and combines it with existing data.
	Combine(req models.StatsRequest)
}

type GetInformation interface {
	IpInformation
	CountryInformation
	CurrencyInformation
	// GetAllProducts retrieves all relevant products for a given IP address.
	GetAllProducts(ip string) error
	// GetStatsService returns an instance of StatsInformation for statistics handling.
	GetStatsService() StatsInformation
}
