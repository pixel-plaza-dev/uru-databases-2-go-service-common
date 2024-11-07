package path

import (
	"os"
	"path/filepath"
)

// GetCurrentDirectory gets the current directory of the executable
func GetCurrentDirectory() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(execPath), nil
}
