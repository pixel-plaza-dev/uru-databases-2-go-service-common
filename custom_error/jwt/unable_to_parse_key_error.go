package jwt

import (
	"github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils"
	"strings"
)

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
	formattedKey := utils.AddParentheses(u.KeyType.String())
	formattedError := utils.AddBrackets(u.Err.Error())
	return strings.Join([]string{"Unable to parse", formattedKey, "key", formattedError}, " ")
}
