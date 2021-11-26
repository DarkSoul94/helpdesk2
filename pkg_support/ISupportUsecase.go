package pkg_support

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type ISupportUsecase interface {
	//Открытие/переоткрытие смены саппорту.
	OpenShift(supportID uint64, user *models.User) models.Err
	//Закрытие смены саппорта. Если смену закрывает администратор - отправляет запросы у этого саппорта на распределение
	CloseShift(supportID uint64, user *models.User) models.Err
	//Получение последней смены по суппорту
	GetLastShift(supportID uint64) (*internal_models.Shift, models.Err)
	//Получить текущий статус саппорта
	GetSupportStatus(supportID uint64) (*internal_models.Status, models.Err)
	//Сменить статус саппорта на указанный
	SetSupportStatus(supportID, statusID uint64) models.Err
	//Получить весь список саппортов
	GetSupportList() ([]*internal_models.Support, models.Err)
	//Получить список саппортов которые принимают запросы в данный момент
	GetActiveSupports() ([]*internal_models.Support, models.Err)
	//Получить список возможных статусов для саппортов
	GetStatusesList() ([]*internal_models.Status, models.Err)
	//Получить информацию о текущем статусе саппортов
	//(статус, состояние смены, количество запросов в разных статуса в разрезе саппортов)
	GetCurrentStatuses() ([]*internal_models.SupportInfo, map[string]int, models.Err)
	//Получить список саппортов которые отмечены в своих карточка как старшие
	GetSeniors() ([]*internal_models.Support, models.Err)

	//Получение карточки саппорта по ID карты
	GetCard(cardID uint64) (*internal_models.Card, models.Err)
	//Получение списка карточек саппортов
	GetCardsList() ([]*internal_models.Card, models.Err)
	//Обновление карточки саппорта
	UpdateCard(card *internal_models.Card) models.Err
}

type ISuppForUser interface {
	//Создание саппортов по ID пользователей
	CreateSupport(usersID ...uint64) models.Err
	//Удаление саппортов по ID пользователей
	DeleteSupport(usersID ...uint64) models.Err
}

type ISuppForTicket interface {
	//Возвращает саппорта на которого нужно распределить запрос.
	//Если свободные саппорты не были найдены возвращает nill
	GetSupportForDistribution(supportID uint64) *internal_models.Support
	//Создает запись о работе саппорта над запросом
	AddSupportActivity(support *internal_models.Support, ticketID uint64) models.Err
	//Удаляет запись о работе саппорта над запросом
	RemoveSupportActivity(ticketID uint64) models.Err
	//Модифицирует запись о работе саппорта над запросом (в случае когда запрос переходит к другому саппорту)
	UpdateSupportActivity(supportID, ticketID uint64) models.Err
}
