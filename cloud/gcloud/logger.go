package gcloud

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
	"google.golang.org/grpc/credentials/oauth"
)

// Logger is the logger for Google Cloud
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for Google Cloud
func NewLogger(logger commonlogger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// LoadedTokenSource logs the loaded token source
func (l *Logger) LoadedTokenSource(tokenSource *oauth.TokenSource) {
	// Check if the token source is nil
	if tokenSource == nil {
		l.logger.LogError(commonlogger.NewLogError("Failed to load token source", NilTokenSourceError))
		return
	}

	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		l.logger.LogError(commonlogger.NewLogError("Failed to load token source", err))
		return
	}

	l.logger.LogMessage(commonlogger.NewLogMessage("Loaded token source", commonlogger.StatusDebug, token.AccessToken))
}
