package logger

type UsersLogger struct {
	logger Logger
}

// NewUsersLogger creates a new users logger
func NewUsersLogger(name string) UsersLogger {
	return UsersLogger{logger: Logger{name}}
}
