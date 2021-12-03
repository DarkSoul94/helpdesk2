package pkg_reports

import "time"

type IReportsRepo interface {
	GetAverageGradesBySupport(startDate, endDate time.Time) (map[string]float64, error)
	Close() error
}
