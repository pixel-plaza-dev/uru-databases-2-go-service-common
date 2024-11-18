package outgoing_ctx

import commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"

// Logger is the logger for the outgoing context debugger
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the outgoing context debugger
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// LogKeyValue logs the key value
func (l Logger) LogKeyValue(key string, value string) {
	formattedKey := "Outgoing context key '" + key + "' value"
	l.logger.LogMessageWithDetails(formattedKey, value)
}
