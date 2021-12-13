package pkg_consts

import "time"

type IConstsRepo interface {
	SetConst(key string, data interface{}) error
	GetConst(key string, target interface{}) error

	CreateHistory(key string, data interface{}) error
	GetHistory(date time.Time, key string, target interface{}) error

	Close() error
}
