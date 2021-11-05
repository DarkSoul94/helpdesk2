package mysql

import "errors"

var (
	ErrReadGroup      = errors.New("Failed read group from db")
	ErrReadGroupsList = errors.New("Failed read users list from db")

	ErrCreateUser    = errors.New("Failed insert new user to db")
	ErrReadUser      = errors.New("Failed read user from db")
	ErrReadUsersList = errors.New("Failed read users list from db")
)
