package pkg_reports

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type IReportsUsecase interface {
	GetTicketStatusDifference(startDate, endDate string) (map[internal_models.TicketDifference][]internal_models.StatusDifference, models.Err)
	GetAverageGradesBySupport(startDate, endDate string) (map[string]float64, models.Err)
	GetTicketsGrade(startDate, endDate string, usersID []uint64, departments []string) (map[string]map[string][]internal_models.TicketGrade, models.Err)
	GetTicketsCountByDaysHours(startDate, endDate string) (map[string]map[string]uint, models.Err)
	GetSupportsStatusesByWeekDay(startDate, endDate string) (map[uint]map[string][]internal_models.SupportStatus, models.Err)
	GetSupportsShifts(startDate, endDate string) (map[string][]internal_models.SupportsShifts, models.Err)
	GetSupportsStatusHistory(date string) (map[string][]internal_models.SupportStatusHistory, models.Err)
}
