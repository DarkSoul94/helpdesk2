package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_reports"
)

type ReportsUsecase struct {
	repo pkg_reports.IReportsRepo
}

func NewReportsUsecase(repo pkg_reports.IReportsRepo) *ReportsUsecase {
	return &ReportsUsecase{
		repo: repo,
	}
}

func (u *ReportsUsecase) GetAverageGradesBySupport(startDate, endDate string) (map[string]float64, models.Err) {
	start, end, err := u.parseTime(startDate, endDate)
	if err != nil {
		return nil, err
	}

	grades, er := u.repo.GetAverageGradesBySupport(start, end)
	if er != nil {
		return nil, models.InternalError(er.Error())
	}

	return grades, nil
}

func (u *ReportsUsecase) parseTime(startDate, endDate string) (time.Time, time.Time, models.Err) {
	var (
		start, end time.Time
		err        error
	)

	if len(startDate) > 0 {
		start, err = time.ParseInLocation(`2006-01-02`, startDate, time.Local)
		if err != nil {
			logger.LogError("Failed parse time", "reports/usecase/", "startDate", err)
			return time.Time{}, time.Time{}, errFailedParseDate
		}
	}

	if len(endDate) > 0 {

		end, err = time.ParseInLocation(`2006-01-02`, endDate, time.Local)
		if err != nil {
			logger.LogError("Failed parse time", "reports/usecase/", "endDate", err)
			return time.Time{}, time.Time{}, errFailedParseDate
		}

		addTime, err := time.ParseDuration(`23h59m59s`)
		if err != nil {
			logger.LogError("Failed parse time", "reports/usecase/", "addTime", err)
			return time.Time{}, time.Time{}, errFailedParseDate
		}
		end = end.Add(addTime)
	}

	return start, end, nil
}
