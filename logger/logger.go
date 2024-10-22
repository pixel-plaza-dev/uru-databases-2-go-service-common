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

// buildMessage creates a string that contains a message that could be either a success or an error
func (l *Logger) buildMessage(message string) string {
	return strings.Join([]string{l.FormattedName, message}, " ")
}

// buildMessageWithDetails creates a string that contains a message with details
func (l *Logger) buildMessageWithDetails(message string, details string) string {
	formattedDetails := utils.AddBrackets(details)
	return strings.Join([]string{l.FormattedName, message, formattedDetails}, " ")
}

// logMessage logs a message
func (l *Logger) logMessage(message string) {
	formattedMessage := l.buildMessage(message)
	log.Println(formattedMessage)
}

// logMessageWithDetails logs a message with details
func (l *Logger) logMessageWithDetails(message string, details string) {
	formattedMessage := l.buildMessageWithDetails(message, details)
	log.Println(formattedMessage)
}
