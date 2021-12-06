package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_reports"
	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type ReportsUsecase struct {
	repo pkg_reports.IReportsRepo
}

func NewReportsUsecase(repo pkg_reports.IReportsRepo) *ReportsUsecase {
	return &ReportsUsecase{
		repo: repo,
	}
}

func (u *ReportsUsecase) GetTicketStatusDifference(startDate, endDate string) (map[internal_models.TicketDifference][]internal_models.StatusDifference, models.Err) {
	start, end, err := u.parseTime(startDate, endDate)
	if err != nil {
		return nil, err
	}

	difference, er := u.repo.GetTicketStatusDifference(start, end)
	if er != nil {
		return nil, models.InternalError(er.Error())
	}

	return difference, nil
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

func (u *ReportsUsecase) GetTicketsGrade(startDate, endDate string, usersID []uint64, departments []string) (map[string]map[string][]internal_models.TicketGrade, models.Err) {
	if len(usersID) == 0 && len(departments) == 0 {
		return nil, models.BadRequest("Не выбраного ни одного пользователя и раздела")
	}

	start, end, err := u.parseTime(startDate, endDate)
	if err != nil {
		return nil, err
	}

	grades, er := u.repo.GetTicketsGrade(start, end, usersID, departments)
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

func (u ReportsUsecase) GetSupportsShifts(startDate, endDate string) (map[string][]internal_models.SupportsShifts, models.Err) {
	var (
		period       internal_models.Period
		rangeByMonth []internal_models.Period
		result       map[string][]internal_models.SupportsShifts = make(map[string][]internal_models.SupportsShifts)
	)

	period, err := internal_models.ParceString(startDate, endDate)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	rangeByMonth = period.SplitByMonth()

	for _, month := range rangeByMonth {
		index := month.FormLabel()

		report, err := u.repo.GetSupportsShifts(month.StartDate, month.EndDate)
		if err != nil {
			return nil, models.InternalError(err.Error())
		}

		result[index] = report
	}

	return result, nil
}

func (u *ReportsUsecase) GetSupportsStatusHistory(date string) (map[string][]internal_models.SupportStatusHistory, models.Err) {
	startDate, endDate, err := u.parseTime(date, date)
	if err != nil {
		return nil, err
	}

	historyList, er := u.repo.GetSupportsStatusHistory(startDate, endDate)
	if er != nil {
		return nil, models.InternalError(er.Error())
	}

	return historyList, nil
}
