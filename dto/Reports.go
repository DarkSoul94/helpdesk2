package dto

import (
	"fmt"
	"sort"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type OutMotivation struct {
	Support           *motivSupp        `json:"support"`
	Categories        []motivCategories `json:"categories"`
	TotalTicketsCount uint64            `json:"total_tickets_count"`
	TotalMotivation   float64           `json:"total_motivation"`
	TotalByShifts     float64           `json:"total_by_shifts"`
	TotalPayment      float64           `json:"total_payment"`
}

func ToOutMotivation(inpMotiv []internal_models.Motivation) []OutMotivation {
	outMotiv := make([]OutMotivation, 0)
	for _, motivation := range inpMotiv {
		totalMotivation, _ := motivation.TotalMotivation.Float64()
		totalShifts, _ := motivation.TotalByShifts.Float64()
		total, _ := motivation.Total.Float64()
		outMotiv = append(outMotiv, OutMotivation{
			Support:           toMotivationSupport(motivation.Support),
			Categories:        toMotivationCategories(motivation.ByCategory),
			TotalMotivation:   totalMotivation,
			TotalByShifts:     totalShifts,
			TotalPayment:      total,
			TotalTicketsCount: motivation.TotalTicketsCount,
		})
	}
	return outMotiv
}

type motivSupp struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func toMotivationSupport(mSupport *internal_models.MotivSupport) *motivSupp {
	return &motivSupp{
		ID:    mSupport.ID,
		Name:  mSupport.Name,
		Color: mSupport.Color,
	}
}

type motivCategories struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	TicketsCount uint64 `json:"tickets_count"`
}

func toMotivationCategories(byCategories []internal_models.MotivCategory) []motivCategories {
	outCategories := make([]motivCategories, 0)
	for _, category := range byCategories {
		outCategories = append(outCategories, motivCategories{
			ID:           category.ID,
			Name:         category.Name,
			TicketsCount: category.Count,
		})
	}
	return outCategories
}

type OutStatusDifference struct {
	Status   string `json:"status"`
	DiffTime string `json:"diff_time"`
}

func toOutStatusDifference(diff internal_models.StatusDifference) OutStatusDifference {
	return OutStatusDifference{
		Status:   diff.StatusName,
		DiffTime: diff.Duration,
	}
}

type OutTicketStatusDifference struct {
	TicketID         uint64                `json:"ticket_id"`
	SupportName      string                `json:"support_name"`
	Section          string                `json:"section"`
	StatusDifference []OutStatusDifference `json:"status_difference"`
}

func ToOutTicketStatusDifference(ticket internal_models.TicketDifference, statuses []internal_models.StatusDifference) OutTicketStatusDifference {
	outTicket := OutTicketStatusDifference{
		TicketID:         ticket.TicketID,
		SupportName:      ticket.SupportName,
		Section:          ticket.Section,
		StatusDifference: make([]OutStatusDifference, 0),
	}

	for _, status := range statuses {
		outTicket.StatusDifference = append(outTicket.StatusDifference, toOutStatusDifference(status))
	}

	return outTicket
}

type InpParam struct {
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	UsersID     []uint64 `json:"users_id,omitempty"`
	Departments []string `json:"departments,omitempty"`
}

type OutTicketGrade struct {
	TicketID    uint64 `json:"ticket_id"`
	TicketGrade uint   `json:"ticket_grade"`
}

//outUserTicketGrades ...
type OutUserTicketGrades struct {
	UserName         string           `json:"user_name"`
	TicketsGrades    []OutTicketGrade `json:"tickets_grades"`
	AverageUserGrade float64          `json:"average_user_grade"`
}

//outDepartmentTicketGrade ...
type OutDepartmentTicketGrade struct {
	Department             string                `json:"department"`
	UsersGrades            []OutUserTicketGrades `json:"users_grades"`
	AvaregeDepartmentGrade float64               `json:"avarege_department_grade"`
}

type OutAverageGrade struct {
	Name         string  `json:"support"`
	AverageGrade float64 `json:"average_grade_by_support"`
}

type OutReturnedTicket struct {
	TicketID    uint64 `json:"ticket_id"`
	TicketDate  string `json:"ticket_date"`
	Category    string `json:"category"`
	Section     string `json:"section"`
	TicketText  string `json:"ticket_text"`
	Status      string `json:"status"`
	Author      string `json:"author"`
	Support     string `json:"support"`
	TicketGrade uint64 `json:"ticket_grade"`
}

func ToOutReturnedTickets(tick internal_models.ReturnedTicket) OutReturnedTicket {
	return OutReturnedTicket{
		TicketID:    tick.TicketID,
		TicketDate:  fmt.Sprint(tick.TicketDate.Format("2006-01-02 15:04:05")),
		Category:    tick.Category,
		Section:     tick.Section,
		TicketText:  tick.TicketText,
		Status:      tick.Status,
		Author:      tick.Author,
		Support:     tick.Support,
		TicketGrade: tick.TicketGrade,
	}
}

type OutTicketsCountByHour struct {
	Hour  string `json:"hour"`
	Count uint   `json:"count"`
}

type OutTicketsCountByDay struct {
	Date        string                  `json:"date"`
	CountByHour []OutTicketsCountByHour `json:"count_by_hour"`
}

func ToOutTicketsCountByDay(counts map[string]map[string]uint) []OutTicketsCountByDay {
	outCounts := make([]OutTicketsCountByDay, 0)

	for day, countByHour := range counts {
		countByDay := OutTicketsCountByDay{
			Date:        day,
			CountByHour: make([]OutTicketsCountByHour, 0),
		}

		for hour, count := range countByHour {
			countByDay.CountByHour = append(countByDay.CountByHour, OutTicketsCountByHour{
				Hour:  hour,
				Count: count,
			})
		}

		sortHour(countByDay.CountByHour)
		outCounts = append(outCounts, countByDay)
	}

	return outCounts
}

func sortHour(countByDay []OutTicketsCountByHour) {
	sort.Slice(countByDay, func(i, j int) bool {
		return countByDay[i].Hour < countByDay[j].Hour
	})
}

type OutSupportStatusesHistory struct {
	WeekDay      uint             `json:"week_day"`
	SupportsList []OutSupportList `json:"supports_list"`
}

type OutSupportList struct {
	SupportName string               `json:"support_name"`
	Statuses    []OutSupportStatuses `json:"statuses"`
}

type OutSupportStatuses struct {
	StatusName string `json:"status_name"`
	Duration   string `json:"duration"`
}

func ToOutSupportStatusesHistory(history map[uint]map[string][]internal_models.SupportStatus) []OutSupportStatusesHistory {
	var outHistory []OutSupportStatusesHistory = make([]OutSupportStatusesHistory, 0)

	for day, supportStatuses := range history {
		outHist := OutSupportStatusesHistory{
			WeekDay:      day,
			SupportsList: make([]OutSupportList, 0),
		}

		for support, statusList := range supportStatuses {
			outSupp := OutSupportList{
				SupportName: support,
				Statuses:    make([]OutSupportStatuses, 0),
			}

			for _, status := range statusList {
				outSupp.Statuses = append(outSupp.Statuses, OutSupportStatuses{
					StatusName: status.StatusName,
					Duration:   (status.Duration * time.Second).String(),
				})
			}

			outHist.SupportsList = append(outHist.SupportsList, outSupp)
		}

		outHistory = append(outHistory, outHist)
	}

	return outHistory
}

type OutSupportsShift struct {
	Support           string              `json:"support"`
	WithOutGraceTime  string              `json:"with_out_grace_time"`
	ShiftsCount       int                 `json:"shifts_count"`
	TotalMinutesCount string              `json:"total_minutes_count"`
	SupportShifts     []outOpeningDayTime `json:"shifts"`
}

type outOpeningDayTime struct {
	OpeningDate        string `json:"opening_date"`
	ClosingDate        string `json:"closing_date"`
	CountOfMinutesLate uint64 `json:"count_of_minutes_late"`
}

func ToOutSupportShift(shift *internal_models.SupportsShifts) OutSupportsShift {
	outShift := OutSupportsShift{
		Support:           shift.Support,
		TotalMinutesCount: time.Duration(shift.MinutesCount * uint64(time.Minute)).String(),
		ShiftsCount:       shift.ShiftsCount,
		WithOutGraceTime:  time.Duration(shift.WithOutGraceTime * uint64(time.Minute)).String(),
	}

	for _, val := range shift.DayTime {
		outShift.SupportShifts = append(outShift.SupportShifts, outOpeningDayTime{
			OpeningDate:        val.OpeningDate,
			ClosingDate:        val.ClosingDate,
			CountOfMinutesLate: val.CountOfMinutesLate,
		})
	}

	return outShift
}

type OutSupportStatusHistory struct {
	StatusName string `json:"time"`
	SelectTime string `json:"name"`
	Duration   uint64 `json:"difference"`
}

type OutSupportStatusHistoryList struct {
	SupportName      string                    `json:"support"`
	StatusesHisttory []OutSupportStatusHistory `json:"statuses"`
}

func ToOutSupportStatusHistoryList(historyList map[string][]internal_models.SupportStatusHistory) []OutSupportStatusHistoryList {
	var outHistoryList []OutSupportStatusHistoryList = make([]OutSupportStatusHistoryList, 0)

	for support, historyList := range historyList {
		supportHistory := OutSupportStatusHistoryList{SupportName: support}

		for _, history := range historyList {
			outHistory := OutSupportStatusHistory{
				StatusName: history.StatusName,
				SelectTime: history.SelectTime.Local().Format("15:04:05"),
			}

			if history.Duration == 0 {
				outHistory.Duration = uint64(time.Now().Local().Sub(history.SelectTime).Seconds())
			} else {
				outHistory.Duration = uint64(history.Duration)
			}

			supportHistory.StatusesHisttory = append(supportHistory.StatusesHisttory, outHistory)
		}

		outHistoryList = append(outHistoryList, supportHistory)
	}

	return outHistoryList
}
