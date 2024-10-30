package jwt

type Token string

const (
	RefreshToken Token = "refresh_token"
	AccessToken  Token = "access_token"
)

// String returns the string representation of the token
func (t Token) String() string {
	return string(t)
}
