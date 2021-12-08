package pkg_reports

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type IReportsRepo interface {
	GetSupportTicketCountByCategory(startDate, endDate time.Time, support_id uint64) (map[uint64]uint64, error)
	GetTicketStatusDifference(startDate, endDate time.Time) (map[internal_models.TicketDifference][]internal_models.StatusDifference, error)
	GetAverageGradesBySupport(startDate, endDate time.Time) (map[string]float64, error)
	GetTicketsGrade(startDate, endDate time.Time, usersID []uint64, departments []string) (map[string]map[string][]internal_models.TicketGrade, error)
	GetReturnedTickets(startDate, endDate time.Time) ([]internal_models.ReturnedTicket, error)
	GetTicketsCountByDaysHours(startDate, endDate time.Time) (map[string]map[string]uint, error)
	GetSupportsStatusesByWeekDay(startDate, endDate time.Time) (map[uint]map[string][]internal_models.SupportStatus, error)
	GetSupportsShifts(startDate, endDate time.Time) ([]internal_models.SupportsShifts, error)
	GetSupportsStatusHistory(startDate, endDate time.Time) (map[string][]internal_models.SupportStatusHistory, error)
	GetConstVal(date time.Time, name string) uint64
	Close() error
}
