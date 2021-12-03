package pkg_reports

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type IReportsRepo interface {
	GetAverageGradesBySupport(startDate, endDate time.Time) (map[string]float64, error)
	GetSupportsShifts(startDate, endDate time.Time) ([]internal_models.SupportsShifts, error)
	GetSupportsStatusHistory(startDate, endDate time.Time) (map[string][]internal_models.SupportStatusHistory, error)
	GetConstVal(date time.Time, name string) uint64
	Close() error
}
