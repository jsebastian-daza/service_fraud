package cmd

import (
	"errors"
	"testing"

	"service_fraud/interfaces"
	"service_fraud/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetInformation struct {
	mock.Mock
}

func (m *MockGetInformation) Geolocation(ip string) models.IpApiResponse {
	args := m.Called(ip)
	return args.Get(0).(models.IpApiResponse)
}

func (m *MockGetInformation) GetCountryInformation(country string) models.CountryResponse {
	args := m.Called(country)
	return args.Get(0).(models.CountryResponse)
}

func (m *MockGetInformation) GetCurrencyInformation() models.CurrencyResponse {
	args := m.Called()
	return args.Get(0).(models.CurrencyResponse)
}

func (m *MockGetInformation) GetAllProducts(ip string) error {
	args := m.Called(ip)
	return args.Error(0)
}

func (m *MockGetInformation) GetStatsService() interfaces.StatsInformation {
	args := m.Called()
	return args.Get(0).(interfaces.StatsInformation)
}

type MockStatsService struct {
	mock.Mock
}

func (m *MockStatsService) GetStats() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockStatsService) Combine(req models.StatsRequest) {
	m.Called(req)
}

// Pruebas unitarias
func TestIsValidIp(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"1.1.1.1", true},
		{"256.256.256.256", false},
		{"invalid_ip", false},
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			err := IsValidIp(tt.ip)

			if tt.expected {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetInformation(t *testing.T) {
	mockService := new(MockGetInformation)

	// Simula un retorno exitoso para GetAllProducts
	mockService.On("GetAllProducts", "1.1.1.1").Return(nil)

	err := GetInformation(mockService, "1.1.1.1")
	assert.NoError(t, err)

	// Simula un retorno con error para GetAllProducts
	mockService.On("GetAllProducts", "2.2.2.2").Return(errors.New("some error"))
	err = GetInformation(mockService, "2.2.2.2")
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
}

func TestStart(t *testing.T) {
	mockGetInformation := new(MockGetInformation)
	mockStatsService := new(MockStatsService)

	// Mockear el comportamiento del servicio de estadísticas
	mockStatsService.On("GetStats").Return("Statistics Data")
	mockGetInformation.On("GetStatsService").Return(mockStatsService)

	getInformationService = mockGetInformation // Asigna el mock al servicio

	t.Run("valid traceip option", func(t *testing.T) {
		mockGetInformation.On("GetAllProducts", "1.1.1.1").Return(nil)

		err := Start("traceip 1.1.1.1")
		assert.NoError(t, err)

	})

	t.Run("valid record option", func(t *testing.T) {
		err := Start("record")
		assert.NoError(t, err)
		mockStatsService.AssertExpectations(t)
	})

	t.Run("invalid option", func(t *testing.T) {
		err := Start("invalid option")
		assert.Error(t, err)
	})
}
