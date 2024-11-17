package flag

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the flag
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// ModeFlagSet is the flag set for the mode
func (l Logger) ModeFlagSet(modeFlag *ModeFlag) {
	l.logger.LogMessageWithDetails("Mode flag set", modeFlag.String())
}
