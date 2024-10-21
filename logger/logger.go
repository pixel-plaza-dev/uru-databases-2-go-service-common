package logger

import "strings"

type Logger struct {
	name string
}

// buildMessage creates a string that contains a message that could be either a success or an error
func (l *Logger) buildMessage(message string) string {
	return l.name + " " + message
}

// buildMessageWithDetails creates a string that contains a message with details
func (l *Logger) buildMessageWithDetails(message string, details string) string {
	return l.name + " " + message + ": " + details
}

// buildErrorMessage creates a string that contains an error message
func (l *Logger) buildErrorMessage(message string, err error) string {
	return strings.Join([]string{l.name, message, err.Error()}, " ")
}
