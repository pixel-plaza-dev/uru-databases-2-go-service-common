package rate_limiter

import "errors"

var (
	TooManyRequestsError = errors.New("too many requests")
)
