package pkg_reports

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type IReportsRepo interface {
	GetAverageGradesBySupport(startDate, endDate time.Time) (map[string]float64, error)
	GetSupportsStatusHistory(startDate, endDate time.Time) (map[string][]internal_models.SupportStatusHistory, error)
	Close() error
}
