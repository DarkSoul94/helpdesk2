package pkg_reports

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type IReportsRepo interface {
	GetTicketStatusDifference(startDate, endDate time.Time) (map[internal_models.TicketDifference][]internal_models.StatusDifference, error)
	GetAverageGradesBySupport(startDate, endDate time.Time) (map[string]float64, error)
	GetTicketsGrade(startDate, endDate time.Time, usersID []uint64, departments []string) (map[string]map[string][]internal_models.TicketGrade, error)
	GetTicketsCountByDaysHours(startDate, endDate time.Time) (map[string]map[string]uint, error)
	GetSupportsShifts(startDate, endDate time.Time) ([]internal_models.SupportsShifts, error)
	GetSupportsStatusHistory(startDate, endDate time.Time) (map[string][]internal_models.SupportStatusHistory, error)
	GetConstVal(date time.Time, name string) uint64
	Close() error
}
