package usecase

import (
	"sync"
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_reports"
	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
)

type ReportsUsecase struct {
	catSecUC  cat_sec_manager.ICatSecUsecase
	scheduler pkg_scheduler.IReportsSchedulerUsecase
	repo      pkg_reports.IReportsRepo
}

func NewReportsUsecase(catSecUC cat_sec_manager.ICatSecUsecase, scheduler pkg_scheduler.IReportsSchedulerUsecase, repo pkg_reports.IReportsRepo) *ReportsUsecase {
	return &ReportsUsecase{
		catSecUC:  catSecUC,
		scheduler: scheduler,
		repo:      repo,
	}
}

func (u *ReportsUsecase) GetMotivation(startDate, endDate string) (map[string][]*internal_models.Motivation, models.Err) {
	var (
		wg               sync.WaitGroup
		mutex            sync.Mutex
		motivByPer              = make(map[string][]*internal_models.Motivation)
		motivByAllPeriod        = make(map[uint64]*internal_models.Motivation)
		key              string = "01. Итого за период"
		supportKeys             = make([]uint64, 0)
	)

	inpPeriod, er := internal_models.ParceString(startDate, endDate)
	if er != nil {
		return nil, models.InternalError(er.Error())
	}

	periods := inpPeriod.SplitByMonth()
	periodCount := len(periods)

	wg.Add(periodCount)

	categoryList, err := u.catSecUC.GetCategoryList()
	if err != nil {
		return nil, err
	}

	for _, period := range periods {
		go func(period internal_models.Period) {
			defer wg.Done()
			index := period.FormLabel()
			//начало расчета мотивации по одному саппорту за интервал времени (до одного месяца)
			shiftMotivation, err := u.scheduler.SupportsShiftsMotivation(period.StartDate, period.EndDate)
			if err != nil {
				return
			}

			for _, shift := range shiftMotivation {

				suppMotiv := internal_models.NewMotivation(
					shift.SupportID,
					shift.SupportName,
					shift.Color,
					shift.Motivation,
				)
				err := u.calculateSupportMotivation(period, suppMotiv, categoryList)
				if err != nil {
					return
				}

				mutex.Lock()
				motivByPer[index] = append(motivByPer[index], suppMotiv)
				mutex.Unlock()

				if periodCount > 1 {
					if _, ok := motivByAllPeriod[shift.SupportID]; !ok {
						supportKeys = append(supportKeys, shift.SupportID)
					}

					internal_models.CalcMotivByAllPeriod(motivByAllPeriod, suppMotiv)
				}
			}
			mutex.Lock()
			motivByPer[index] = append(motivByPer[index], internal_models.MotivByPeriod(motivByPer[index]))
			mutex.Unlock()
		}(period)
	}

	wg.Wait()

	if periodCount > 1 {
		for _, id := range supportKeys {
			motivByPer[key] = append(motivByPer[key], motivByAllPeriod[id])
		}

		motivByPer[key] = append(motivByPer[key], internal_models.MotivByPeriod(motivByPer[key]))
	}

	return motivByPer, nil
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

func (u *ReportsUsecase) GetTicketsCountByDaysHours(startDate, endDate string) (map[string]map[string]uint, models.Err) {
	start, end, err := u.parseTime(startDate, endDate)
	if err != nil {
		return nil, err
	}
	list, er := u.repo.GetTicketsCountByDaysHours(start, end)
	if er != nil {
		return nil, models.InternalError(er.Error())
	}

	return list, nil
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

func (u *ReportsUsecase) GetReturnedTickets(startDate, endDate string) ([]internal_models.ReturnedTicket, models.Err) {
	start, end, err := u.parseTime(startDate, endDate)
	if err != nil {
		return nil, err
	}

	tickets, er := u.repo.GetReturnedTickets(start, end)
	if er != nil {
		return nil, models.InternalError(er.Error())
	}

	return tickets, nil
}

func (u *ReportsUsecase) GetSupportsStatusesByWeekDay(startDate, endDate string) (map[uint]map[string][]internal_models.SupportStatus, models.Err) {
	start, end, err := u.parseTime(startDate, endDate)
	if err != nil {
		return nil, err
	}
	history, er := u.repo.GetSupportsStatusesByWeekDay(start, end)
	if er != nil {
		return nil, models.InternalError(er.Error())
	}

	return history, nil
}

func (u ReportsUsecase) GetSupportsShifts(startDate, endDate string) (map[string][]*internal_models.SupportsShifts, models.Err) {
	var (
		period       internal_models.Period
		rangeByMonth []internal_models.Period
		result       map[string][]*internal_models.SupportsShifts = make(map[string][]*internal_models.SupportsShifts)
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
		completeSupportShifts(report)
		result[index] = report
	}

	return result, nil
}

func completeSupportShifts(list []*internal_models.SupportsShifts) {
	for _, val := range list {
		val.ShiftsCount = len(val.DayTime)
		for _, shift := range val.DayTime {
			val.MinutesCount += shift.CountOfMinutesLate
		}
	}
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
