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
	case pkg_consts.DBKeyBanner:
		if val, ok := data["text"]; ok {
			if err := u.repo.SetConst(key, val); err != nil {
				return models.InternalError("Ошибка заполнения новым значением")
			}

			return nil
		} else {
			return models.BadRequest("Поле данных 'text' пустое")
		}
	default:
		return models.BadRequest(fmt.Sprintf("Константы '%s' не существует", key))
	}
}

func (u *ConstsUsecase) GetConst(key string) (map[string]interface{}, models.Err) {
	data := make(map[string]interface{})

	switch key {
	case pkg_consts.DBKeyBanner:
		var text string

		err := u.repo.GetConst(key, &text)
		if err != nil {
			return nil, models.InternalError(fmt.Sprintf("Не удалось получить значение константы '%s'", key))
		}

		data["text"] = text
	default:
		return nil, models.BadRequest(fmt.Sprintf("Константы '%s' не существует", key))
	}

	return data, nil
}
