package pkg_scheduler

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
)

type ISchedulerRepo interface {
	//Записывает в базу новый офис
	AddOffice(office *internal_models.Office) models.Err
	//Обновляет в базе новый офис
	UpdateOffice(office *internal_models.Office) models.Err
	//Получение офиса из базы по его ID
	GetOfficeByID(officeID uint64) (*internal_models.Office, models.Err)
	//Получение списка офисов. Признак `deleted` определяет какие офисы нужно достать из базы - удаленные или нет.
	//Если в функцию также передать дату, то она вернет список офисов которые были указаны в графике за эту дату
	GetOfficesList(deleted bool, dates ...string) ([]*internal_models.Office, models.Err)

	//Создает или обновляет ячейку графика.
	//Если ячейка была создана то в объект ячейки записывается присвоенный ей ID
	UpdateCell(cell *internal_models.Cell) models.Err
	//Принимает в себя хэш-карту в которой ключи это месяцы за которые обновляется график,
	//а значения по ключам - ID ячеек которые должны остаться в графике.
	//Функция удаляет по каждому месяцу все ячейки которых нет в массивах ID-шников
	DeleteCells(actualCellsIDs map[string][]uint64) models.Err
	//Получает массив ячеек графика за определенную дату
	GetSchedule(date string) ([]*internal_models.Cell, models.Err)
	//Получить ячейку с графиком по саппорту за сегодня
	GetTodayShift(supportID uint64) (*internal_models.Cell, models.Err)
	//Получает количество смен по всем суппортам
	GetShiftsCount(startDate, endDate time.Time) (map[uint64]int64, models.Err)

	//Создание запизаписи об опоздании
	CreateLateness(lateness *internal_models.Lateness) models.Err
	//Получение списка опозданий за указанный месяц
	GetLateness(date string) ([]*internal_models.Lateness, models.Err)
	//Получение конкретной записи об опоздании по ее ID
	GetLatenessByID(latenessId uint64) (*internal_models.Lateness, models.Err)
	//Обновление записи об опоздании
	UpdateLateness(lateness *internal_models.Lateness) models.Err
	//Получить наличие новых записей об опоздании без решений
	CheckNewLateness() bool

	Close()
}
