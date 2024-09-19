package models

import (
	"fmt"
	"service_fraud/utils"
	"strconv"
	"strings"
	"time"
)

// Response represents the structured output for IP, country, and currency information.
type Response struct {
	Ip                string
	CurrentDate       time.Time
	Country           string
	ISO               string
	Lenguages         string
	Currency          string
	CurrentTime       time.Time
	EstimatedDistance string
}

// FormatResponse formats and displays the response based on provided data.
func (r *Response) FormatResponse(ipRes IpApiResponse, countryRes CountryResponse, currencyRes CurrencyResponse) {
	currencyArr := make([]string, 0, len(countryRes.ArrayResponse[0].Currencies))
	var str string
	for k, _ := range countryRes.ArrayResponse[0].Currencies {
		currencyArr = append(currencyArr, k)
	}

	str = fmt.Sprintf(`
		IP: %s,  fecha actual: %s
			País: %s
			ISO Code: %s`,
		ipRes.IP, time.Now().Format("2006-01-02 15:04:05"),
		ipRes.CountryName,
		countryRes.ArrayResponse[0].Cca2,
	)

	for _, v := range ipRes.Location.Languages {
		str += fmt.Sprintf("\n			Idiomas: %s (%s)", v.Name, v.Code)
	}

	str += fmt.Sprintf(`
			Moneda: %v (1 %v = %f U$S)`,
		currencyArr[0], currencyArr[0], getCurrencyRates(currencyArr[0], currencyRes),
	)

	for i, _ := range countryRes.ArrayResponse[0].Timezones {
		hour, _ := parseOffset(countryRes.ArrayResponse[0].Timezones[i])
		str += fmt.Sprintf("\n			Hora: %s (UTC) o %s (%s)", time.Now().UTC().Format("2006-01-02 15:04:05"), hour, countryRes.ArrayResponse[0].Timezones[i])
	}

	str += fmt.Sprintf(`
			Distancia Estimada: %s kms (%f, %f) a (%f, %f)
	`, utils.GetEstimatedDistance(utils.BA_LATITUDE, utils.BA_LONGITUDE, ipRes.Latitude, ipRes.Longitude), utils.BA_LATITUDE, utils.BA_LONGITUDE, ipRes.Latitude, ipRes.Longitude)

	fmt.Print(str)
}

// getCurrencyRates calculates the exchange rate for the requested currency in terms of USD.
func getCurrencyRates(requestedCurrency string, rates CurrencyResponse) float64 {
	usdRate := rates.Rates["USD"]
	requestedRate := rates.Rates[requestedCurrency]
	return (1.0 * ((1 / requestedRate) * usdRate))
}

// parseOffset converts a timezone offset string to a formatted local time.
func parseOffset(offsetStr string) (string, error) {
	// Verificar si el offset es válido
	if !strings.HasPrefix(offsetStr, "UTC") {
		return "", fmt.Errorf("offset debe comenzar con 'UTC'")
	}

	// Extraer el signo y el tiempo
	offsetStr = strings.TrimPrefix(offsetStr, "UTC")
	sign := offsetStr[:1]
	offsetStr = offsetStr[1:]

	// Analizar las horas y minutos
	parts := strings.Split(offsetStr, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("offset inválido")
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("error al analizar las horas: %v", err)
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("error al analizar los minutos: %v", err)
	}

	// Calcular el desplazamiento en segundos
	totalMinutes := hours*60 + minutes
	if sign == "-" {
		totalMinutes = -totalMinutes
	}
	offsetSeconds := totalMinutes * 60

	// Crear una zona horaria fija con el offset calculado
	location := time.FixedZone(offsetStr, offsetSeconds)

	now := time.Now().UTC()

	// Convertir la hora actual a la zona horaria especificada
	localTime := now.In(location)

	return localTime.Format("2006-01-02 15:04:05"), nil
}
