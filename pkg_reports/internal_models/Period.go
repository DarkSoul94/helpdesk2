package internal_models

import (
	"fmt"
	"time"
)

type Period struct {
	StartDate time.Time
	EndDate   time.Time
}

func ParceString(start, end string) (Period, error) {
	var (
		err error
		p   Period
	)
	p.StartDate, err = time.Parse("2006-01-02", start)
	if err != nil {
		return Period{}, err
	}
	p.EndDate, err = time.Parse("2006-01-02", end)
	if err != nil {
		return Period{}, err
	}

	return p, nil
}

func (p *Period) SplitByMonth() []Period {
	var (
		periodList []Period = make([]Period, 0)
	)

	startDate := p.StartDate
	endDate := p.EndDate

	if equal(startDate, endDate) {
		periodList = append(periodList, formPeriod(startDate, endDate))
		return periodList
	}
	periodList = append(periodList, formPeriod(startDate, monthLastDay(startDate)))

	startDate = monthFirstDay(startDate)
	for {
		startDate = startDate.AddDate(0, 1, 0)
		if equal(startDate, endDate) {
			periodList = append(periodList, formPeriod(startDate, endDate))
			break
		} else {
			periodList = append(periodList, formPeriod(startDate, monthLastDay(startDate)))
		}
	}
	return periodList
}

func formPeriod(start, end time.Time) Period {
	return Period{
		StartDate: start,
		EndDate:   addTime(end, "23h59m59s"),
	}
}

func addTime(date time.Time, t string) time.Time {
	addTime, _ := time.ParseDuration(t)

	return date.Add(addTime)
}

func monthFirstDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

func monthLastDay(date time.Time) time.Time {
	return monthFirstDay(date).AddDate(0, 1, -1)
}

func equal(start, end time.Time) bool {
	return start.Year() == end.Year() && start.Month() == end.Month()
}

func (p *Period) FormLabel() string {
	return fmt.Sprintf("%s ~ %s", p.StartDate.Format("2006-01-02"), p.EndDate.Format("2006-01-02"))
}
