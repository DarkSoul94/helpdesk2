package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
)

type OutScheduleCell struct {
	ID        uint64 `json:"id"`
	SupportID uint64 `json:"support_id"`
	OfficeID  uint64 `json:"office_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Date      string `json:"date"`
	Vacation  bool   `json:"vacation"`
	SickLeave bool   `json:"sick_leave"`
}

func ToModelScheduleCell(cell OutScheduleCell) *internal_models.Cell {
	date, _ := time.Parse(`2006-01-02`, cell.Date)
	start, _ := time.Parse(`15:04`, cell.StartTime)
	end, _ := time.Parse(`15:04`, cell.EndTime)
	return &internal_models.Cell{
		ID:        cell.ID,
		SupportID: cell.SupportID,
		OfficeID:  cell.OfficeID,
		Date:      date,
		StartTime: start,
		EndTime:   end,
		Vacation:  cell.Vacation,
		SickLeave: cell.SickLeave,
	}
}
