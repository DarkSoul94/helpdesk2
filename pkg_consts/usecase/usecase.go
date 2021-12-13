package usecase

import (
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_consts"
	"github.com/shopspring/decimal"
)

type ConstsUsecase struct {
	repo pkg_consts.IConstsRepo
}

func NewConstsUsecase(repo pkg_consts.IConstsRepo) *ConstsUsecase {
	return &ConstsUsecase{
		repo: repo,
	}
}

func (u *ConstsUsecase) SetConst(key string, data map[string]interface{}) models.Err {
	switch key {
	case pkg_consts.KeyBanner:
		return u.setBanner(data)

	case pkg_consts.KeyConfig:
		return u.setSettings(data)

	default:
		return models.BadRequest(fmt.Sprintf("Константы '%s' не существует", key))
	}
}

func (u *ConstsUsecase) GetConst(key string) (map[string]interface{}, models.Err) {

	switch key {
	case pkg_consts.KeyBanner:
		return u.getBanner()

	case pkg_consts.KeyConfig:
		return u.getSettings()

	default:
		return nil, models.BadRequest(fmt.Sprintf("Константы '%s' не существует", key))
	}
}

func (u *ConstsUsecase) GetConstForDate(date time.Time, key string) (map[string]interface{}, models.Err) {
	switch key {
	case pkg_consts.KeyConfig:
		var (
			graceTime uint64
			penalty   decimal.Decimal
		)

		data := u.getSettingsHistory(date)
		if _, ok := data[key_Grace]; !ok {
			u.repo.GetConst(key_Grace, &graceTime)
			data[key_Grace] = graceTime
		}
		if _, ok := data[key_Penalty]; !ok {
			u.repo.GetConst(key_Penalty, &penalty)
			data[key_Penalty] = penalty
		}
		return data, nil

	default:
		return nil, models.BadRequest(fmt.Sprintf("Для константы '%s' не предусмотрено сохранение истории изменений", key))
	}
}
