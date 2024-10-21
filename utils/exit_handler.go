package utils

import (
	"os"
)

type Exit struct{ Code int }

// ExitHandler is a defer function that handles the exit of the program
func ExitHandler() {
	// Recover from panic
	if err := recover(); err != nil {
		// Check if the error is an Exit
		if exit, ok := err.(Exit); ok == true {
			os.Exit(exit.Code)
		}

		// Not an Exit, re-panic
		panic(err)
	}
}
