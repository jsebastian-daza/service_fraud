package utils

import (
	"math"
	"strconv"
)

// Constants for user messages, error messages, and API URLs used in the application.
const (
	INFO_USER_MESSAGE_SELECT_OPTION     = "Ingrese su opcion: "
	NO_RECORD_INFORMATION_AVAILABLE_YET = "Aun no hay informacion disponible para visualizar"
	ERR_USER_MESSAGE_INVALID_OPTION     = "Opcion invalida, por favor revise la informacion ingresada"
	ERR_MESSAGE_INVALID_OPTION          = "Invalid option. Input: %s"
	ERR_CODE_INVALID_OPTION             = 101
	ERR_USER_MESSAGE_INVALID_IP         = "La ip ingresada no es valida, revise los datos ingresados e intente nuevamente"
	ERR_MESSAGE_INVALID_IP              = "The IP is not valid: %s"
	ERR_CODE_INVALID_IP                 = 102
	ERR_USER_MESSAGE_IP_SERVICE         = "Error durante la ejecucion de la obtencion de la ip"
	ERR_MESSAGE_IP_SERVICE              = "Error during the execution of the IP service: %s"
	ERR_CODE_IP_SERVICE                 = 103
	ERR_USER_MESSAGE_COUNTRY_SERVICE    = "Error durante la ejecucion de la obtencion de la informacion de la region"
	ERR_MESSAGE_COUNTRY_SERVICE         = "Error during the execution of the country service: %s"
	ERR_CODE_COUNTRY_SERVICE            = 104
	ERR_USER_MESSAGE_CURRENCY_SERVICE   = "Error durante la ejecucion de la obtencion de la informacion de la moneda"
	ERR_MESSAGE_CURRENCY_SERVICE        = "Error during the execution of the currency service: %s"
	ERR_CODE_CURRENCY_SERVICE           = 105
	ERR_USER_MESSAGE_GET_SECRETS        = "Error durante la ejecucion de la obtencion de los secretos"
	ERR_MESSAGE_GET_SECRETS             = "Error getting the secrets: %s"
	ERR_CODE_GET_SECRETS                = 106
	ERR_USER_MESSAGE_IP_RESP_EMPTY      = "La ip solicitada no devuelve informacion para mostrar"
	ERR_MESSAGE_IP_RESP_EMPTY           = "The requested IP returned an empty response: %s"
	ERR_CODE_IP_RESP_EMPTY              = 107
	ERR_USER_MESSAGE_LIMIT_REACHED      = "El servicio alcanzo su limite permitido es necesario generar una nueva clase"
	ERR_MESSAGE_LIMIT_REACHED           = "It is necessary to create a new api key: %s"
	ERR_CODE_LIMIT_REACHED              = 108

	LOG_MESSAGE_VALID_PARAMETER = "Opcion valida iniciando el proceso para: %s"
	LOG_MESSAGE_ELAPSED_TIME    = "Tiempo transcurrido para el flujo %s: %f (segundos)"

	API_IP_URL       = "http://api.ipapi.com/api/%s?access_key=%s"
	API_COUNTRY_URL  = "https://restcountries.com/v3.1/name/%s?fullText=true"
	API_CURRENCY_URL = "https://data.fixer.io/api/latest?access_key=%s"

	SECRET_VAULT            = "service_fraud_api_secrets"
	SECRET_API_IP_KEY       = "ipapi_key"
	SECRET_API_CURRENCY_KEY = "currency_key"

	BA_LATITUDE  float64 = -34.61315
	BA_LONGITUDE float64 = -58.37723
	EARTH_RADIUS float64 = 6371.0

	TTL_IN_MINUTES = 30
)

// ToRadians converts degrees to radians.
func ToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// GetEstimatedDistance calculates the distance between two geographic points
// (given in latitude and longitude) using the Haversine formula and returns it as a string.
func GetEstimatedDistance(lat1, lon1, lat2, lon2 float64) string {
	lat1 = ToRadians(lat1)
	lon1 = ToRadians(lon1)
	lat2 = ToRadians(lat2)
	lon2 = ToRadians(lon2)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EARTH_RADIUS * c
	return strconv.Itoa(int(distance))
}
