package flag

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the logger for the flag
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the flag
func NewLogger(logger commonlogger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// ModeFlagSet is the flag set for the mode
func (l Logger) ModeFlagSet(mode *ModeFlag) {
	// Check if the mode flag is nil
	if mode == nil {
		return
	}

	l.logger.LogMessage(commonlogger.NewLogMessage("Mode flag set", commonlogger.StatusDebug, mode.String()))
}
