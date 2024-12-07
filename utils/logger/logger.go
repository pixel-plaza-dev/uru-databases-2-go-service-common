package logger

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"log"
	"strings"
)

type (
	// Log interface
	Log interface {
		String() string
	}

	// LogMessage struct
	LogMessage struct {
		Title   string
		Details []string
		Status  Status
	}

	// LogError struct
	LogError struct {
		Title  string
		Errors []error
	}

	// Logger is an interface for logging messages
	Logger interface {
		LogMessage(logMessage *LogMessage)
		LogError(logError *LogError)
	}

	// DefaultLogger is a logger that logs messages
	DefaultLogger struct {
		Name          string
		FormattedName string
	}
)

// FormatDetails gets the formatted details
func FormatDetails(details []string) string {
	return utils.AddBrackets(strings.Join(details, ", "))
}

// FormatStatus gets the formatted status
func FormatStatus(status Status) string {
	return utils.AddBrackets(status.String())
}

// NewLogMessage creates a new log message
func NewLogMessage(title string, status Status, details ...string) *LogMessage {
	return &LogMessage{Title: title, Status: status, Details: details}
}

// String gets the string representation of a log message
func (l *LogMessage) String() string {
	var formattedLog []string

	// Format status
	if l.Status != StatusNone {
		formattedLog = append(formattedLog, FormatStatus(l.Status))
	}

	// Add title
	if l.Title != "" {
		formattedLog = append(formattedLog, l.Title)
	}

	// Add formatted details
	if len(l.Details) > 0 {
		formattedLog = append(formattedLog, FormatDetails(l.Details))
	}

	return strings.Join(formattedLog, " ")
}

// NewLogError creates a new log error
func NewLogError(title string, errors ...error) *LogError {
	return &LogError{Title: title, Errors: errors}
}

// FormatErrors gets the formatted errors
func FormatErrors(errors []error) string {
	var errorsString []string
	for _, err := range errors {
		errorsString = append(errorsString, err.Error())
	}

	// Get formatted errors
	return utils.AddBrackets(strings.Join(errorsString, ", "))
}

// String gets the string representation of a log error
func (l *LogError) String() string {
	var formattedLog []string

	// Add message
	if l.Title != "" {
		formattedLog = append(formattedLog, l.Title)
	}

	// Add formatted errors
	if len(l.Errors) > 0 {
		formattedLog = append(formattedLog, FormatErrors(l.Errors))
	}

	return strings.Join(formattedLog, " ")
}

// NewDefaultLogger creates a new logger
func NewDefaultLogger(name string) DefaultLogger {
	return DefaultLogger{Name: name, FormattedName: utils.AddBrackets(name)}
}

// FormatLogMessage formats a log message
func (d DefaultLogger) FormatLogMessage(logMessage *LogMessage) string {
	if logMessage == nil {
		return d.FormattedName
	}

	return strings.Join([]string{d.FormattedName, logMessage.String()}, " - ")
}

// LogMessage logs a message
func (d DefaultLogger) LogMessage(logMessage *LogMessage) {
	log.Println(d.FormatLogMessage(logMessage))
}

// FormatLogError formats a log error
func (d DefaultLogger) FormatLogError(logError *LogError) string {
	if logError == nil {
		return d.FormattedName
	}

	return strings.Join(
		[]string{
			d.FormattedName,
			FormatStatus(StatusFailed),
			logError.String(),
		}, " - ",
	)
}

// LogError logs an error
func (d DefaultLogger) LogError(logError *LogError) {
	log.Println(d.FormatLogError(logError))
}
