package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"
)

type OutAverageGrade struct {
	Name         string  `json:"support"`
	AverageGrade float64 `json:"average_grade_by_support"`
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
