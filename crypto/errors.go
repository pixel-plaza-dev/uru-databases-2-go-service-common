package crypto

import "errors"

var FailedToHashPasswordError = errors.New("failed to hash password")
var PasswordNotHashedError = errors.New("password is not hashed")
