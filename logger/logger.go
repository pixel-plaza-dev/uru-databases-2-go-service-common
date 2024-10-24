package logger

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"log"
	"strings"
)

type Logger struct {
	Name          string
	FormattedName string
}

// NewLogger creates a new logger
func NewLogger(name string) *Logger {
	return &Logger{Name: name, FormattedName: utils.AddBrackets(name)}

}

// BuildMessage creates a string that contains a message that could be either a success or an error
func (l *Logger) BuildMessage(message string) string {
	return strings.Join([]string{l.FormattedName, message}, " ")
}

// BuildMessageWithDetails creates a string that contains a message with details
func (l *Logger) BuildMessageWithDetails(message string, details string) string {
	formattedDetails := utils.AddBrackets(details)
	return strings.Join([]string{l.FormattedName, message, formattedDetails}, " ")
}

// LogMessage logs a message
func (l *Logger) LogMessage(message string) {
	formattedMessage := l.BuildMessage(message)
	log.Println(formattedMessage)
}

// LogMessageWithDetails logs a message with details
func (l *Logger) LogMessageWithDetails(message string, details string) {
	formattedMessage := l.BuildMessageWithDetails(message, details)
	log.Println(formattedMessage)
}
