package usecase

import (
	"fmt"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/shopspring/decimal"
)

func (u *ConstsUsecase) getBanner() (map[string]interface{}, models.Err) {
	data := make(map[string]interface{})
	var text string
	err := u.repo.GetConst(key_Banner, &text)
	if err != nil {
		return nil, models.InternalError(fmt.Sprintf("Не удалось получить значение константы '%s'", key_Banner))
	}
	data["text"] = text
	return data, nil
}

func (u *ConstsUsecase) setBanner(data map[string]interface{}) models.Err {
	if val, ok := data["text"]; ok {
		if err := u.repo.SetConst(key_Banner, val); err != nil {
			return models.InternalError("Ошибка заполнения новым значением")
		}
		return nil
	}
	return models.BadRequest("Поле данных 'text' пустое")
}

func (u *ConstsUsecase) getSettings() (map[string]interface{}, models.Err) {
	data := make(map[string]interface{})
	var (
		penalty   decimal.Decimal
		graceTime uint64
	)
	if err := u.repo.GetConst(key_Penalty, &penalty); err != nil {
		data[key_Penalty], _ = decimal.Zero.Float64()
	}
	data[key_Penalty] = penalty

	if err := u.repo.GetConst(key_Grace, &graceTime); err != nil {
		data[key_Grace] = 0
	}
	data[key_Grace] = graceTime

	return data, nil
}

func (u *ConstsUsecase) setSettings(data map[string]interface{}) models.Err {
	var (
		penalty   decimal.Decimal
		graceTime uint64
	)

	if val, ok := data[key_Penalty]; ok {
		u.repo.GetConst(key_Penalty, &penalty)
		if val != penalty {
			if err := u.repo.SetConst(key_Penalty, decimal.NewFromFloat(val.(float64))); err != nil {
				return models.InternalError("Ошибка заполнения новым значением")
			}
			u.repo.CreateHistory(key_Penalty, penalty)
		}
	}

	if val, ok := data[key_Grace]; ok {
		u.repo.GetConst(key_Grace, &graceTime)
		if val != graceTime {
			if err := u.repo.SetConst(key_Grace, uint64(val.(float64))); err != nil {
				return models.InternalError("Ошибка заполнения новым значением")
			}

			u.repo.CreateHistory(key_Grace, graceTime)
		}
	}

	return nil
}
