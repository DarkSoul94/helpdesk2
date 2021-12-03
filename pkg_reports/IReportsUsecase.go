package pkg_reports

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type IReportsUsecase interface {
	GetAverageGradesBySupport(startDate, endDate string) (map[string]float64, models.Err)
	GetSupportsStatusHistory(date string) (map[string][]internal_models.SupportStatusHistory, models.Err)
}
