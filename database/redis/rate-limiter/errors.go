package rate_limiter

import "errors"

var (
	TooManyRequestsError = errors.New("Too many requests")
)
