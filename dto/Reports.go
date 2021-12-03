package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

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
