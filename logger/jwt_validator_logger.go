package logger

type JwtValidatorLogger struct {
	logger *Logger
}

// NewJwtValidatorLogger creates a new JWT validator logger
func NewJwtValidatorLogger(name string) *JwtValidatorLogger {
	return &JwtValidatorLogger{logger: NewLogger(name)}
}

// ValidatedToken logs a message when the server validates a token
func (e *JwtIssuerLogger) ValidatedToken() {
	e.logger.logMessage("Validated token")
}
