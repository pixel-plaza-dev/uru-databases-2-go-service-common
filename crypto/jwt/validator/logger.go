package validator

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the JWT validator logger
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger creates a new JWT validator logger
func NewLogger(logger commonlogger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// ValidatedToken logs a message when the server validates a token
func (l Logger) ValidatedToken() {
	l.logger.LogMessage(
		commonlogger.NewLogMessage(
			"Validated token",
			commonlogger.StatusInfo,
		),
	)
}

// MissingTokenClaimsUserId logs the missing token claims user ID
func (l Logger) MissingTokenClaimsUserId() {
	l.logger.LogMessage(
		commonlogger.NewLogMessage(
			"Missing  user ID in token claims",
			commonlogger.StatusFailed,
		),
	)
}
