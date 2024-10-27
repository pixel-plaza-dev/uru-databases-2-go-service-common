package flag

import "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"

type Logger struct {
	logger *logger.Logger
}

// NewLogger is the logger for the flag
func NewLogger(name string) *Logger {
	return &Logger{logger: logger.NewLogger(name)}
}

// ModeFlagSet is the flag set for the mode
func (l *Logger) ModeFlagSet(modeFlag *ModeFlag) {
	l.logger.LogMessageWithDetails("Mode flag set", modeFlag.String())
}
