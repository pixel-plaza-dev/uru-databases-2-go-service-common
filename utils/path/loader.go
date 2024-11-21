package path

import (
	"os"
)

// ReadFile reads the file from the given path
func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, UnableToReadKeyFileError
	}
	return file, nil
}
