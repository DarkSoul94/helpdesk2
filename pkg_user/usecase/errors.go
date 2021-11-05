package usecase

import "errors"

var (
	ErrUnauthorized = errors.New("You are not authorized to perform this action")

	ErrFailedGetUsersList = errors.New("Failed get users list")

	ErrFailedGetGroupList = errors.New("Failed get groups list")
)
