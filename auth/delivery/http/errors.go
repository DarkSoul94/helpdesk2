package http

import "errors"

var (
	ErrInvalidAccessToken   = errors.New("Invalid access token")
	ErrUserNameIsBlank      = errors.New("Username is blank")
	ErrPassIsBlank          = errors.New("Password is blank")
	ErrPassRequirements     = errors.New("Password does not meet security requirements")
	ErrUserNameRequirements = errors.New("Username does not meet requirements")
)
