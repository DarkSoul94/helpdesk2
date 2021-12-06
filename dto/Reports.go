package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

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

type outOpeningDayTime struct {
	OpeningDate        string `json:"opening_date"`
	ClosingDate        string `json:"closing_date"`
	CountOfMinutesLate uint64 `json:"count_of_minutes_late"`
}

type OutSupportsShift struct {
	Support          string              `json:"support"`
	WithOutGraceTime string              `json:"with_out_grace_time"`
	SupportShifts    []outOpeningDayTime `json:"shifts"`
}

func ToOutSupportShift(shift internal_models.SupportsShifts) OutSupportsShift {
	outShift := OutSupportsShift{
		Support:          shift.Support,
		WithOutGraceTime: time.Duration(shift.WithOutGraceTime * uint64(time.Minute)).String(),
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
