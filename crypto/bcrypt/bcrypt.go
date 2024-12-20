package bcrypt

import (
	"errors"
	commoncrypto "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return "", commoncrypto.FailedToHashPasswordError
	}

	return string(hash), nil
}

// CheckPasswordHash checks if the password matches the hash
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// IsHashed checks if the password is hashed
// IsHashed checks if a string is a bcrypt hash
func IsHashed(str string) bool {
	// bcrypt hashes are always 60 characters long
	if len(str) != 60 {
		return false
	}

	// Try to decode the hash
	err := bcrypt.CompareHashAndPassword([]byte(str), []byte{})
	if err != nil {
		return errors.Is(err, bcrypt.ErrMismatchedHashAndPassword)
	}
	return true
}
