package info

import (
	"strings"
)

// GetMethodName gets the method name from the full method
func GetMethodName(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}
