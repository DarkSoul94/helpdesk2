package pkg_reports

import "github.com/DarkSoul94/helpdesk2/models"

type IReportsUsecase interface {
	GetAverageGradesBySupport(startDate, endDate string) (map[string]float64, models.Err)
}
