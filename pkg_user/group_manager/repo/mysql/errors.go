package mysql

import "errors"

var (
	ErrReadGroup      = errors.New("Failed read group from db")
	ErrReadGroupsList = errors.New("Failed read users list from db")
)
