package validator

type DuplicateKeyError struct {
	Field string
	Key   string
}

// Error returns the error message
func (e DuplicateKeyError) Error() string {
	return "Unique field '" + e.Field + "' with key '" + e.Key + "' already exists"
}
