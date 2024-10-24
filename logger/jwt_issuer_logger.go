package logger

type JwtIssuerLogger struct {
	logger *Logger
}

// NewJwtIssuerLogger creates a new JWT issuer logger
func NewJwtIssuerLogger(name string) *JwtIssuerLogger {
	return &JwtIssuerLogger{logger: NewLogger(name)}
}

// IssuedToken logs a message when the server issues a token
func (e *JwtIssuerLogger) IssuedToken() {
	e.logger.LogMessage("Issued token")
}
