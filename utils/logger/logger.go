package logger

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"log"
	"strings"
)

type (
	// Logger is an interface for logging messages
	Logger interface {
		BuildMessage(message string) string
		BuildMessageWithDetails(message string, details string) string
		BuildMessageWithDetailsAndStatus(
			message string, details string, status Status) string
		LogMessage(message string)
		LogMessageWithDetails(message string, details string)
		LogMessageWithDetailsAndStatus(
			message string, details string, status Status)
	}

	// DefaultLogger is a logger that logs messages
	DefaultLogger struct {
		Name          string
		FormattedName string
	}
)

// NewDefaultLogger creates a new logger
func NewDefaultLogger(name string) DefaultLogger {
	return DefaultLogger{Name: name, FormattedName: utils.AddBrackets(name)}
}

// BuildMessage creates a string that contains a message that could be either a success or an error
func (d DefaultLogger) BuildMessage(message string) string {
	return strings.Join([]string{d.FormattedName, message}, " ")
}

// BuildMessageWithDetails creates a string that contains a message with details
func (d DefaultLogger) BuildMessageWithDetails(
	message string, details string,
) string {
	formattedDetails := utils.AddBrackets(details)
	return strings.Join(
		[]string{d.FormattedName, message, formattedDetails}, " ",
	)
}

// BuildMessageWithDetailsAndStatus creates a string that contains a message with details and a status
func (d DefaultLogger) BuildMessageWithDetailsAndStatus(
	message string, details string, status Status,
) string {
	formattedDetails := utils.AddBrackets(details)
	formattedStatus := utils.AddBrackets(status.String())
	return strings.Join(
		[]string{d.FormattedName, message, formattedDetails, formattedStatus}, " ",
	)
}

// LogMessage logs a message
func (d DefaultLogger) LogMessage(message string) {
	formattedMessage := d.BuildMessage(message)
	log.Println(formattedMessage)
}

// LogMessageWithDetails logs a message with details
func (d DefaultLogger) LogMessageWithDetails(message string, details string) {
	formattedMessage := d.BuildMessageWithDetails(message, details)
	log.Println(formattedMessage)
}

// LogMessageWithDetailsAndStatus logs a message with details and a status
func (d DefaultLogger) LogMessageWithDetailsAndStatus(
	message string, details string, status Status,
) {
	formattedMessage := d.BuildMessageWithDetailsAndStatus(message, details, status)
	log.Println(formattedMessage)
}
