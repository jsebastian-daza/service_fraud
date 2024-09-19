package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service_fraud/interfaces"
	"service_fraud/models"
	"service_fraud/utils"
	"time"
)

// InformationService provides methods to retrieve information about IPs, countries, and currencies.
type InformationService struct {
	secrets           interfaces.SecretsVault
	StatsService      interfaces.StatsInformation
	processed         chan models.StatsRequest
	countryDataStore  interfaces.DataStore[string, models.CountryResponse]
	currencyDataStore interfaces.DataStore[string, models.CurrencyResponse]
}

// NewInformationService creates a new instance of InformationService.
func NewInformationService(secrets interfaces.SecretsVault,
	countryDs interfaces.DataStore[string, models.CountryResponse],
	currencyDs interfaces.DataStore[string, models.CurrencyResponse]) *InformationService {

	processed := make(chan models.StatsRequest, WorkerCount)
	statService := NewStatsService(processed)
	return &InformationService{
		secrets:           secrets,
		StatsService:      statService,
		processed:         processed,
		countryDataStore:  countryDs,
		currencyDataStore: currencyDs,
	}
}

// Geolocation fetches geolocation information for a given IP address.
func (s *InformationService) Geolocation(ip string) models.IpApiResponse {
	value, err := s.secrets.GetSecret(utils.SECRET_API_IP_KEY)
	ipresp := models.IpApiResponse{}
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_GET_SECRETS, err)
		ipresp.Error = *models.NewErrorIpApiError(utils.ERR_CODE_GET_SECRETS, utils.ERR_USER_MESSAGE_GET_SECRETS)
		return ipresp
	}
	url := fmt.Sprintf(utils.API_IP_URL, ip, *value)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_IP_SERVICE, err)
		ipresp.Error = *models.NewErrorIpApiError(utils.ERR_CODE_IP_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_IP_SERVICE))
		return ipresp
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_IP_SERVICE, err)
		ipresp.Error = *models.NewErrorIpApiError(utils.ERR_CODE_IP_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_IP_SERVICE))
		return ipresp
	}

	if resp.StatusCode != http.StatusOK {
		str := fmt.Sprintf("Error getting the request: %s", resp.Status)
		log.Println(str)
		ipresp.Error = *models.NewErrorIpApiError(utils.ERR_CODE_IP_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_IP_SERVICE))
		return ipresp
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&ipresp); err != nil {
		log.Printf("error: can't decode - %s \n", err)
		ipresp.Error = *models.NewErrorIpApiError(utils.ERR_CODE_IP_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_IP_SERVICE))
		return ipresp
	}
	return ipresp
}

// GetCountryInformation fetches information for a given country.
func (s *InformationService) GetCountryInformation(country string) models.CountryResponse {
	countryresp := models.CountryResponse{}
	arr := &countryresp.ArrayResponse
	url := fmt.Sprintf(utils.API_COUNTRY_URL, country)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_COUNTRY_SERVICE, err)
		countryresp.Error = *models.NewCountryApiError(utils.ERR_CODE_COUNTRY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_COUNTRY_SERVICE))
		return countryresp
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_COUNTRY_SERVICE, err)
		countryresp.Error = *models.NewCountryApiError(utils.ERR_CODE_COUNTRY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_COUNTRY_SERVICE))
		return countryresp
	}

	if resp.StatusCode != http.StatusOK {
		str := fmt.Sprintf("Error getting the request: %s", resp.Status)
		log.Println(str)
		countryresp.Error = *models.NewCountryApiError(utils.ERR_CODE_COUNTRY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_COUNTRY_SERVICE))
		return countryresp
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&arr); err != nil {
		log.Printf("error: can't decode - %s \n", err)
		countryresp.Error = *models.NewCountryApiError(utils.ERR_CODE_COUNTRY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_COUNTRY_SERVICE))
		return countryresp
	}

	return countryresp
}

// GetCurrencyInformation fetches current currency information.
func (s *InformationService) GetCurrencyInformation() models.CurrencyResponse {
	value, err := s.secrets.GetSecret(utils.SECRET_API_CURRENCY_KEY)
	currencyResponse := models.CurrencyResponse{}
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_GET_SECRETS, err)
		currencyResponse.Error = *models.NewCurrencyApiError(utils.ERR_CODE_GET_SECRETS, fmt.Sprintf(utils.ERR_USER_MESSAGE_IP_SERVICE))
		return currencyResponse
	}

	url := fmt.Sprintf(utils.API_CURRENCY_URL, *value)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_CURRENCY_SERVICE, err)
		currencyResponse.Error = *models.NewCurrencyApiError(utils.ERR_CODE_CURRENCY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_CURRENCY_SERVICE))
		return currencyResponse
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf(utils.ERR_MESSAGE_CURRENCY_SERVICE, err)
		currencyResponse.Error = *models.NewCurrencyApiError(utils.ERR_CODE_CURRENCY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_CURRENCY_SERVICE))
		return currencyResponse
	}

	if resp.StatusCode != http.StatusOK {
		str := fmt.Sprintf("Error getting the request: %s", resp.Status)
		log.Println(str)
		currencyResponse.Error = *models.NewCurrencyApiError(utils.ERR_CODE_CURRENCY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_CURRENCY_SERVICE))
		return currencyResponse
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&currencyResponse); err != nil {
		log.Printf("error: can't decode - %s \n", err)
		currencyResponse.Error = *models.NewCurrencyApiError(utils.ERR_CODE_CURRENCY_SERVICE, fmt.Sprint(utils.ERR_USER_MESSAGE_CURRENCY_SERVICE))
		return currencyResponse
	}
	return currencyResponse
}

// GetStatsService returns the stats service instance.
func (s *InformationService) GetStatsService() interfaces.StatsInformation {
	return s.StatsService
}

// GetAllProducts processes all information related to products based on an IP address.
func (s *InformationService) GetAllProducts(ip string) error {
	response := models.Response{}
	ipResponse := s.Geolocation(ip)
	if ipResponse.HasError() {
		return &ipResponse.Error
	}
	if !ipResponse.ContainsValidResponse() {
		log.Printf(utils.ERR_MESSAGE_IP_RESP_EMPTY, ip)
		return models.NewErrorIpApiError(utils.ERR_CODE_IP_RESP_EMPTY, utils.ERR_USER_MESSAGE_IP_RESP_EMPTY)
	}

	stats := models.StatsRequest{
		Country: ipResponse.RegionName,
		Lat:     ipResponse.Latitude,
		Lon:     ipResponse.Longitude,
	}

	s.processed <- stats

	countryResponse, err := s.countryDataStore.Get(ipResponse.RegionName)
	if err != nil {
		countryResponse = s.GetCountryInformation(ipResponse.CountryName)
		if countryResponse.HasError() {
			return &countryResponse.Error
		}
		s.countryDataStore.Set(ipResponse.RegionName, countryResponse)
	}

	currencyResponse, err := s.currencyDataStore.Get("currency")
	if err != nil {
		currencyResponse = s.GetCurrencyInformation()
		if currencyResponse.HasError() {
			if currencyResponse.Error.Message == "" {
				currencyResponse.Error.Message = utils.ERR_USER_MESSAGE_LIMIT_REACHED
			}
			return &currencyResponse.Error
		}
	}

	response.FormatResponse(ipResponse, countryResponse, currencyResponse)
	return nil
}
