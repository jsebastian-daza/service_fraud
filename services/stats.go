package services

import (
	"fmt"
	"log"
	"service_fraud/models"
	"service_fraud/utils"
	"sort"
	"strconv"
	"sync"
)

const WorkerCount = 3

var instanceStats *StatsService
var onceCreationStats sync.Once

// InformationService proporciona una implementación para obtener información de IP, país y moneda.
type StatsService struct {
	lock             sync.Mutex
	StatsRecord      []models.Stats
	processedChannel chan models.StatsRequest
	StatsChannel     chan models.Stats
	Done             chan struct{}
	Result           string
}

// NewStatsService initializes the StatsService and starts worker goroutines for processing stats.
func NewStatsService(processedChannel chan models.StatsRequest) *StatsService {
	onceCreationStats.Do(func() {
		instanceStats = &StatsService{
			processedChannel: processedChannel,
		}
		for i := 0; i < WorkerCount; i++ {
			go instanceStats.processStats()
		}
	})
	return instanceStats
}

// processStats listens for incoming stats requests and processes them.
func (s *StatsService) processStats() {
	log.Println("Start stats processing started")
	for {
		select {
		case stats, ok := <-s.processedChannel:
			if !ok {
				fmt.Println("Processed channel closed")
				return
			}
			s.Combine(stats)
		case <-s.Done:
			fmt.Println("The stats service is stopped")
			return
		}
	}
}

// GetStats returns the current stats result or a message if none are available.
func (s *StatsService) GetStats() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Result == "" {
		return utils.NO_RECORD_INFORMATION_AVAILABLE_YET
	} else {
		return s.Result
	}

}

// Combine processes a stats request and updates the stats record.
func (s *StatsService) Combine(req models.StatsRequest) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.validateStatsRecord(req)
}

// validateStatsRecord checks if a country already has an entry in the stats record and updates it.
func (s *StatsService) validateStatsRecord(req models.StatsRequest) {
	found := false

	for i := 0; i < len(s.StatsRecord); i++ {
		stats := s.StatsRecord[i]
		if stats.Country == req.Country {
			found = true
			s.StatsRecord[i].Invokes++
		}
	}

	if !found {
		s.StatsRecord = append(s.StatsRecord, models.Stats{
			Country:  req.Country,
			Distance: utils.GetEstimatedDistance(utils.BA_LATITUDE, utils.BA_LONGITUDE, req.Lat, req.Lon),
			Invokes:  1,
		})
	}

	s.Result = fmt.Sprintln("\n" + s.getLowerDistance() + "\n" + s.getHigherDistance() + "\n" + s.getAverageDistance())
}

// getLowerDistance finds and returns the country with the closest distance to Buenos Aires.
func (s *StatsService) getLowerDistance() string {
	var lowerDistance models.Stats
	isSet := false // Bandera para verificar si lowerDistance ha sido asignado

	for i := 0; i < len(s.StatsRecord); i++ {
		if !isSet {
			lowerDistance = s.StatsRecord[i]
			isSet = true
		} else {
			curDist, _ := strconv.ParseFloat(lowerDistance.Distance, 64)
			arrDist, _ := strconv.ParseFloat(s.StatsRecord[i].Distance, 64)
			if arrDist < curDist {
				lowerDistance = s.StatsRecord[i]
			}
		}
	}

	return fmt.Sprintf("Distancia más cercana a Buenos Aires consultada: \n %s con una distancia aproximada de: %s kms",
		lowerDistance.Country, lowerDistance.Distance)
}

// getHigherDistance finds and returns the country with the furthest distance from Buenos Aires.
func (s *StatsService) getHigherDistance() string {
	var higherDistance models.Stats
	isSet := false // Bandera para verificar si higherDistance ha sido asignado

	for i := 0; i < len(s.StatsRecord); i++ {
		if !isSet {
			higherDistance = s.StatsRecord[i]
			isSet = true
		} else {
			curDist, _ := strconv.ParseFloat(higherDistance.Distance, 64)
			arrDist, _ := strconv.ParseFloat(s.StatsRecord[i].Distance, 64)
			if arrDist > curDist {
				higherDistance = s.StatsRecord[i]
			}
		}
	}

	return fmt.Sprintf("Distancia más lejana a Buenos Aires consultada: \n %s con una distancia aproximada de: %s kms",
		higherDistance.Country, higherDistance.Distance)
}

// getAverageDistance calculates and returns the average distance of all recorded countries.
func (s *StatsService) getAverageDistance() string {
	average := 0
	totalInvokes := 0
	s.StatsRecord = orderStats(s.StatsRecord)
	var str string = "==============================\n"
	for i := 0; i < len(s.StatsRecord); i++ {
		stats := s.StatsRecord[i]
		intDistance, _ := strconv.Atoi(stats.Distance)
		totalInvokes += stats.Invokes
		average += (stats.Invokes * intDistance)
		str += fmt.Sprintf(`%s -- %s (kms) -- %d invocaciones %s`, stats.Country, stats.Distance, stats.Invokes, "\n")
	}
	str += "==============================\n"

	str += fmt.Sprintf("Distancia promedio entre las peticiones : %d (kms)", (average / totalInvokes))

	return str
}

// orderStats sorts the stats records by country name.
func orderStats(stats []models.Stats) []models.Stats {
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Country < stats[j].Country
	})

	return stats
}
