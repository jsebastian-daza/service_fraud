package utils

import (
	"fmt"
	"log"
	"os"
)

// init sets up logging for the application by creating or opening
// a file named "app.log". Log messages will be appended to this file.
func init() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
	fmt.Print("Log setup")
}
