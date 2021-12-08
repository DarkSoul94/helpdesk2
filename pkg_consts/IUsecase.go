package pkg_consts

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
)

type IConstsUsecase interface {
	SetConst(key string, data map[string]interface{}) models.Err
	GetConst(key string) (map[string]interface{}, models.Err)
	GetConstForDate(date time.Time, key string) (map[string]interface{}, models.Err)
}
