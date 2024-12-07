package validator

import (
	"github.com/golang-jwt/jwt/v5"
	pbtypesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Validator does parsing and validation of JWT token
type (
	Validator interface {
		GetToken(tokenString string) (*jwt.Token, error)
		GetClaims(tokenString string) (*jwt.MapClaims, error)
		GetValidatedClaims(
			token string,
			interception pbtypesgrpc.Interception,
		) (*jwt.MapClaims, error)
		ValidateTokenType() error
	}
)
