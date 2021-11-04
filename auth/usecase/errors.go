package usecase

import "errors"

var (
	ErrLoginFailed = errors.New("Invalid user name or password")
	ErrCreateToken = errors.New("Fail to create tocken for user") 
)
