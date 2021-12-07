package usecase

import (
	"fmt"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_consts"
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
