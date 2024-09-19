// main.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"service_fraud/cmd"
	"service_fraud/models"
	"service_fraud/utils"
	"strings"
)

// main is the entry point of the application. It displays a welcome message,
// continuously processes user input until "exit" is entered, and handles errors.
func main() {
	fmt.Println(cmd.Logo)
	fmt.Println(utils.INFO_USER_MESSAGE_SELECT_OPTION)

	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			os.Exit(0)
		}

		log.Println("Start Process")
		err := cmd.Start(input)
		if err != nil {
			handleError(err)
		}
		fmt.Println("\n" + utils.INFO_USER_MESSAGE_SELECT_OPTION)
	}
}

// handleError processes the provided error and displays an appropriate message
// based on the type of error encountered, such as IpApiError, CountryApiError,
// or CurrencyApiError, providing specific feedback to the user.
func handleError(err error) {
	apiError := &models.IpApiError{}
	countryError := &models.CountryApiError{}
	currencyError := &models.CurrencyApiError{}

	switch {
	case errors.As(err, &apiError):
		fmt.Println(apiError.Error())
	case errors.As(err, &countryError):
		fmt.Println(countryError.Error())
	case errors.As(err, &currencyError):
		fmt.Println(currencyError.Error())
	default:
		fmt.Println(utils.ERR_USER_MESSAGE_INVALID_OPTION)
	}
}
