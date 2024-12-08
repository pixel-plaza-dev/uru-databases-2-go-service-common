package outgoing_ctx

import commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"

// Logger is the logger for the outgoing context debugger
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the outgoing context debugger
func NewLogger(logger commonlogger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// LogKeyValue logs the key value
func (l *Logger) LogKeyValue(key string, value string) {
	formattedKey := "Outgoing context key '" + key + "' value"
	l.logger.LogMessage(commonlogger.NewLogMessage(formattedKey, commonlogger.StatusDebug, value))
}
