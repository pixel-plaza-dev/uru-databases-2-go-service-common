package jwt

type KeyType string

const (
	PrivateKey KeyType = "private"
	PublicKey  KeyType = "public"
)

// String returns a string representation of the KeyType
func (k KeyType) String() (keyType string) {
	return string(k)
}

type UnableToParseKeyError struct {
	KeyType KeyType
	Err     error
}

// Error returns a formatted error message for UnableToParseKeyError
func (u UnableToParseKeyError) Error() (message string) {
	return "Unable to parse '" + u.KeyType.String() + "' key: " + u.Err.Error()
}
