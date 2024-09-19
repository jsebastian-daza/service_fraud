package cmd

import (
	"fmt"
	"log"
	"net"
	"service_fraud/interfaces"
	"service_fraud/models"
	"service_fraud/services"
	"service_fraud/utils"
	"strings"
	"time"
)

// Logo is the ASCII art logo displayed in the application.
var Logo = `	
__   _   ___   __  _   _   _   ___   __  _   _   __     __                       
|  \ | | | __| |  \| | | \ / | | __| |  \| | | | | _\   /__\                      
| -< | | | _|  | | ' |  \ V /  | _|  | | ' | | | | v | | \/ |                     
|__/ |_|_|___| |_|\__|_  \_/   |___| |_|\__| |_|_|__/___\__/  _   _    ___   __   
| _,\ | _ \ | || | | __| |  \  /  \   |_   _| | __|  / _/ |  \| | | |  / _/  /  \  
| v_/ | v / | \/ | | _|  | -< | /\ |    | |   | _|  | \__ | | ' | | | | \__ | /\ | 
|_| __|_|_\_ \__/  |___| |__/ |_||_|    |_|   |___|  \__/ |_|\__| |_|  \__/ |_||_| 
|  V  | | __| | |   | |                                                           
| \_/ | | _|  | |_  | |                                                           
|_| |_| |___| |___| |_| 

RESPUESTA ANTI FRAUDES

Para el funcionamiento del proceso por favor escriba alguna de las siguientes 
opciones

- 'traceip <IP a consultar>' para iniciar el proceso de busqueda. Ejemplo:
 traceip 1.4.193.15

- 'record' para mostrar el resumen y detalle de los registros realizados

- 'exit' salir del programa`

var getInformationService interfaces.GetInformation
var countryRequestDataStore interfaces.DataStore[string, models.CountryResponse]
var currencyRequestDataStore interfaces.DataStore[string, models.CurrencyResponse]

// init initializes the data stores and information service used in the application.
func init() {
	countryRequestDataStore = services.NewRequestDataStore[string, models.CountryResponse]()
	currencyRequestDataStore = services.NewRequestDataStore[string, models.CurrencyResponse]()
	getInformationService = services.NewInformationService(services.NewAwsSecrets(), countryRequestDataStore, currencyRequestDataStore)
}

// Start processes the user option, validates it, and either retrieves information
// about an IP address or provides statistics based on the selected flow.
func Start(option string) error {
	flow, ip, err := isValidOption(option)
	if err != nil {
		log.Println("Error", err.Error())
		return err
	}
	log.Printf(utils.LOG_MESSAGE_VALID_PARAMETER, option)
	if flow == 2 {
		since := time.Now()
		fmt.Print(getInformationService.GetStatsService().GetStats())
		log.Printf(utils.LOG_MESSAGE_ELAPSED_TIME, option, time.Since(since).Seconds())
	} else if flow == 1 {
		since := time.Now()
		err := GetInformation(getInformationService, ip)
		log.Printf(utils.LOG_MESSAGE_ELAPSED_TIME, option, time.Since(since).Seconds())
		return err
	}
	return nil
}

// isValidOption validates the user input option and returns the flow type, IP address
// (if applicable), and any errors encountered during validation.
func isValidOption(option string) (int, string, error) {
	santizeStr := strings.TrimSpace(option)
	if len(santizeStr) == 0 {
		return 0, "", models.NewOptionInvalidError(utils.ERR_CODE_INVALID_OPTION, fmt.Sprintf(utils.ERR_MESSAGE_INVALID_OPTION, option))
	}

	arr := strings.Split(santizeStr, " ")
	switch num := len(arr); {
	case num == 2:
		if arr[0] == "traceip" {
			err := IsValidIp(arr[1])
			if err != nil {
				return 0, "", err
			}
			return 1, arr[1], nil
		} else {
			return 0, arr[1], models.NewOptionInvalidError(utils.ERR_CODE_INVALID_OPTION, fmt.Sprintf(utils.ERR_MESSAGE_INVALID_OPTION, option))
		}
	case num == 1:
		if arr[0] == "record" {
			return 2, "", nil
		} else {
			return 0, "", models.NewOptionInvalidError(utils.ERR_CODE_INVALID_OPTION, fmt.Sprintf(utils.ERR_MESSAGE_INVALID_OPTION, option))
		}
	default:
		return 0, "", models.NewOptionInvalidError(utils.ERR_CODE_INVALID_OPTION, fmt.Sprintf(utils.ERR_MESSAGE_INVALID_OPTION, option))
	}
}

// IsValidIp checks if the provided string is a valid IPv4 address.
// Returns an error if the IP is invalid.
func IsValidIp(ipStr string) error {
	ip := net.ParseIP(ipStr)
	res := ip != nil && ip.To4() != nil
	if !res {
		return models.NewOptionInvalidError(utils.ERR_CODE_INVALID_IP, fmt.Sprintf(utils.ERR_MESSAGE_INVALID_IP, ipStr))
	}
	return nil
}

// GetInformation retrieves all product information for the specified IP address
// using the provided process interface.
func GetInformation(process interfaces.GetInformation, ip string) error {
	err := process.GetAllProducts(ip)
	return err
}
