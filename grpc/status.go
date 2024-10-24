package grpc

type Status int

const (
	Failed Status = iota
	Success
)

// StatusBool returns a boolean value for the status
func (s Status) StatusBool() bool {
	return s == Success
}
