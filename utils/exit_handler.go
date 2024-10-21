package utils

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
	"os"
)

type Exit struct{ Code int }

// ExitHandler is a defer function that handles the exit of the program
func ExitHandler(logger *logger.ListenerLogger) {
	// Recover from panic
	if err := recover(); err != nil {
		// Check if the custom_error is an Exit
		if exit, ok := err.(Exit); ok == true {
			os.Exit(exit.Code)
		}

		// Not an Exit, re-panic
		panic(err)
	}
}
