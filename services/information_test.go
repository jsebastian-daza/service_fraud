package services

import (
	"errors"
	"service_fraud/models"
	"service_fraud/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSecretsVault struct {
	mock.Mock
}

func (m *MockSecretsVault) GetSecret(name string) (*string, error) {
	args := m.Called(name)
	return args.Get(0).(*string), args.Error(1)
}

type MockDataStoreCountry struct {
	mock.Mock
}

func (m *MockDataStoreCountry) Set(key string, value models.CountryResponse) error {
	return m.Called(key, value).Error(0)
}

func (m *MockDataStoreCountry) Get(key string) (models.CountryResponse, error) {
	args := m.Called(key)
	return args.Get(0).(models.CountryResponse), args.Error(1)
}

func (m *MockDataStoreCountry) Expire(key string) error {
	return m.Called(key).Error(0)
}

type MockDataStoreCurrency struct {
	mock.Mock
}

func (m *MockDataStoreCurrency) Set(key string, value models.CurrencyResponse) error {
	return m.Called(key, value).Error(0)
}

func (m *MockDataStoreCurrency) Get(key string) (models.CurrencyResponse, error) {
	args := m.Called(key)
	return args.Get(0).(models.CurrencyResponse), args.Error(1)
}

func (m *MockDataStoreCurrency) Expire(key string) error {
	return m.Called(key).Error(0)
}

func TestGeolocation_Success(t *testing.T) {
	mockSecrets := new(MockSecretsVault)
	apiKey := "dummyApiKey"
	mockSecrets.On("GetSecret", utils.SECRET_API_IP_KEY).Return(&apiKey, nil)

	service := InformationService{secrets: mockSecrets}
	ipResponse := service.Geolocation("192.168.1.1")

	assert.NotNil(t, ipResponse)

}

func TestGetCountryInformation_Success(t *testing.T) {
	mockSecrets := new(MockSecretsVault)
	mockCountryStore := new(MockDataStoreCountry)
	service := NewInformationService(mockSecrets, mockCountryStore, nil)

	mockCountryStore.On("Get", "CountryName").Return(models.CountryResponse{}, nil)

	countryResponse := service.GetCountryInformation("CountryName")

	assert.NotNil(t, countryResponse)
}

func TestGetCountryInformation_ErrorOnRequest(t *testing.T) {
	mockSecrets := new(MockSecretsVault)
	mockCountryStore := new(MockDataStoreCountry)
	service := NewInformationService(mockSecrets, mockCountryStore, nil)

	mockCountryStore.On("Get", "CountryName").Return(models.CountryResponse{}, errors.New("not found"))

	countryResponse := service.GetCountryInformation("CountryName")

	assert.True(t, countryResponse.HasError())
}

func TestGetCurrencyInformation_Success(t *testing.T) {
	mockSecrets := new(MockSecretsVault)
	apiKey := "dummyCurrencyApiKey"
	mockSecrets.On("GetSecret", utils.SECRET_API_CURRENCY_KEY).Return(&apiKey, nil)

	service := InformationService{secrets: mockSecrets}
	currencyResponse := service.GetCurrencyInformation()

	assert.NotNil(t, currencyResponse)
}
