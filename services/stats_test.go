package services

import (
	"testing"

	"service_fraud/models"
	"service_fraud/utils"

	"github.com/stretchr/testify/assert"
)

func TestStatsService_NewStatsService(t *testing.T) {
	processedChannel := make(chan models.StatsRequest)

	statsService := NewStatsService(processedChannel)

	assert.NotNil(t, statsService)
}

func TestStatsService_GetStats_NoRecords(t *testing.T) {
	processedChannel := make(chan models.StatsRequest)
	statsService := NewStatsService(processedChannel)

	result := statsService.GetStats()
	assert.Equal(t, utils.NO_RECORD_INFORMATION_AVAILABLE_YET, result)
}

func TestStatsService_Combine_NewCountry(t *testing.T) {

	processedChannel := make(chan models.StatsRequest)
	statsService := NewStatsService(processedChannel)

	req := models.StatsRequest{
		Country: "Colombia",
		Lat:     -34.61,
		Lon:     -58.38,
	}
	statsService.Combine(req)

	assert.Equal(t, 1, len(statsService.StatsRecord))
	assert.Equal(t, "Colombia", statsService.StatsRecord[0].Country)
	assert.Equal(t, 1, statsService.StatsRecord[0].Invokes)
}

func TestStatsService_Combine_ExistingCountry(t *testing.T) {
	processedChannel := make(chan models.StatsRequest)
	statsService := NewStatsService(processedChannel)
	statsService.StatsRecord = make([]models.Stats, 0)

	req1 := models.StatsRequest{
		Country: "Argentina",
		Lat:     -34.61,
		Lon:     -58.38,
	}
	statsService.Combine(req1)

	req2 := models.StatsRequest{
		Country: "Argentina",
		Lat:     -34.61,
		Lon:     -58.38,
	}

	statsService.Combine(req2)

	assert.Equal(t, 1, len(statsService.StatsRecord))
	assert.Equal(t, 2, statsService.StatsRecord[0].Invokes)
}

func TestStatsService_GetStats_WithRecords(t *testing.T) {
	processedChannel := make(chan models.StatsRequest)
	statsService := NewStatsService(processedChannel)

	req := models.StatsRequest{
		Country: "Argentina",
		Lat:     -34.61,
		Lon:     -58.38,
	}
	statsService.Combine(req)

	result := statsService.GetStats()

	assert.Contains(t, result, "Distancia más cercana a Buenos Aires")
	assert.Contains(t, result, "Distancia más lejana a Buenos Aires")
	assert.Contains(t, result, "Distancia promedio entre las peticiones")
}
