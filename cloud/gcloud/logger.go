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
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// LoadedTokenSource logs the loaded token source
func (l Logger) LoadedTokenSource(tokenSource *oauth.TokenSource) {
	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		l.logger.LogMessageWithDetails(
			"Failed to load token source",
			err.Error(),
		)
		return
	}

	l.logger.LogMessageWithDetails("Loaded token source", token.AccessToken)
}
