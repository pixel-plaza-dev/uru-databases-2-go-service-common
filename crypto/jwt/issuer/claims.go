package issuer

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt"
	"time"
)

// GetExpirationTime returns the expiration time for the given duration
func GetExpirationTime(
	issuedTime time.Time,
	duration time.Duration,
) time.Time {
	return issuedTime.Add(duration)
}

// GenerateClaims generates a new claims object
func GenerateClaims(
	jwtId string,
	userId string,
	userUUID uuid.UUID,
	issuedTime time.Time,
	expirationTime time.Time,
	isRefreshToken bool,
) *jwt.MapClaims {
	return &jwt.MapClaims{
		"exp":                         expirationTime.Unix(),
		"iat":                         issuedTime.Unix(),
		commonjwt.IdClaim:             jwtId,
		commonjwt.UserIdClaim:         userId,
		commonjwt.UserSharedIdClaim:   userUUID.String(),
		commonjwt.IsRefreshTokenClaim: isRefreshToken,
	}
}
