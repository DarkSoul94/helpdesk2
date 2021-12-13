package mysql

import (
	"database/sql"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
)

//toDbShift ...
func (r *SchedulerRepo) toDbOffice(mShift *internal_models.Office) dbOffice {
	return dbOffice{
		ID:      mShift.ID,
		Name:    mShift.Name,
		Color:   mShift.Color,
		Deleted: mShift.Deleted,
	}
}

//toModelOffice ...
func (r *SchedulerRepo) toModelOffice(dbOffice dbOffice) *internal_models.Office {
	return &internal_models.Office{
		ID:      dbOffice.ID,
		Name:    dbOffice.Name,
		Color:   dbOffice.Color,
		Deleted: dbOffice.Deleted,
	}
}

//toDbShiftsScheduleCell ...
func (r *SchedulerRepo) toDbShiftsScheduleCell(cell *internal_models.Cell) dbCell {
	var office sql.NullInt64
	start := cell.StartTime.Format("15:04:05")
	end := cell.EndTime.Format("15:04:05")
	if cell.OfficeID != 0 {
		office = sql.NullInt64{
			Valid: true,
			Int64: int64(cell.OfficeID),
		}
	} else {
		office = sql.NullInt64{
			Valid: false,
		}
	}
	return dbCell{
		ID:        cell.ID,
		SupportID: cell.SupportID,
		OfficeID:  office,
		Date:      cell.Date,
		StartTime: start,
		EndTime:   end,
		Vacation:  cell.Vacation,
		SickLeave: cell.SickLeave,
	}
}

//toModelShiftsScheduleCell ...
func (r *SchedulerRepo) toModelShiftsScheduleCell(cell dbCell) *internal_models.Cell {
	var officeID int64
	start, _ := time.Parse("15:04:05", cell.StartTime)
	end, _ := time.Parse("15:04:05", cell.EndTime)
	if cell.OfficeID.Valid {
		officeID = cell.OfficeID.Int64
	} else {
		officeID = 0
	}
	return &internal_models.Cell{
		ID:        cell.ID,
		SupportID: cell.SupportID,
		OfficeID:  uint64(officeID),
		StartTime: start,
		EndTime:   end,
		Date:      cell.Date,
		Vacation:  cell.Vacation,
		SickLeave: cell.SickLeave,
	}
}

func (r *SchedulerRepo) toModelSupportLateness(dbLate *dbLateness) *internal_models.Lateness {
	decision := internal_models.SetLateDecision(dbLate.Decision.Valid, dbLate.Decision.Bool)
	out := &internal_models.Lateness{
		ID:          dbLate.ID,
		Date:        dbLate.Date,
		SupportID:   dbLate.SupportID,
		SupportName: dbLate.SupportName,
		Cause:       dbLate.Cause,
		Difference:  dbLate.Difference,
		Decision:    &decision,
	}
	return out
}

func (r *SchedulerRepo) toDbLateness(lateness *internal_models.Lateness) dbLateness {
	out := dbLateness{
		ID:         lateness.ID,
		Date:       lateness.Date,
		SupportID:  lateness.SupportID,
		Cause:      lateness.Cause,
		Difference: uint64(lateness.Difference),
	}
	valid, value := lateness.Decision.GetDecisionValue()
	out.Decision = sql.NullBool{
		Valid: valid,
		Bool:  value,
	}
	return out
}
