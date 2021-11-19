package pkg_support

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type ISupportUsecase interface {
	//Создание саппортов по ID пользователей
	CreateSupport(usersID ...uint64) models.Err
	//Удаление саппортов по ID пользователей
	DeleteSupport(usersID ...uint64) models.Err
	//Возвращает саппорта на которого нужно распределить запрос.
	//Если свободные саппорты не были найдены возвращает nill
	GetSupportForDistribution(supportID uint64) *internal_models.Support

	AddSupportActivity(support *internal_models.Support, ticketID uint64) models.Err
	RemoveSupportActivity(ticketID uint64) models.Err
	UpdateSupportActivity(supportID, ticketID uint64) models.Err

	GetSupportList() ([]*internal_models.Support, models.Err)
	GetActiveSupports() ([]*internal_models.Support, models.Err)
	GetStatusesList() ([]*internal_models.Status, models.Err)
	SetSupportStatus(supportID, statusID uint64) models.Err
}
