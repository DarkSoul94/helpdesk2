package pkg_consts

type IConstsRepo interface {
	SetConst(key string, data interface{}) error
	GetConst(key string, target interface{}) error
	CreateHistory(key string, data interface{}) error

	Close() error
}
