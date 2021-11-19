package mysql

import (
	"database/sql"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

func (r *Repo) toDbSupport(support *internal_models.Support) dbSupport {
	return dbSupport{
		SupportID: support.ID,
		StatusID:  support.Status.ID,
		Priority:  support.Priority,
	}
}

func (r *Repo) toModelSupport(dbSupp *dbSupport) *internal_models.Support {
	return &internal_models.Support{
		ID:   dbSupp.SupportID,
		Name: dbSupp.Name,
		Status: &internal_models.Status{
			ID: dbSupp.StatusID,
		},
		Priority: dbSupp.Priority,
	}
}

func (r *Repo) toModelsStatus(status *dbStatus) *internal_models.Status {
	return &internal_models.Status{
		ID:           status.ID,
		Name:         status.Name,
		AcceptTicket: status.AcceptTicket,
	}
}

func (r *Repo) toModelShift(shift *dbShift) *internal_models.Shift {
	mShift := &internal_models.Shift{
		ID: shift.ID,
		Support: &internal_models.Support{
			ID: shift.SupportID,
		},
		OpeningTime:   shift.OpeningTime,
		ClosingStatus: shift.ClosingStatus,
	}
	if shift.ClosingTime.Valid {
		mShift.ClosingTime = shift.ClosingTime.Time
	} else {
		mShift.ClosingTime = time.Time{}
	}
	return mShift
}

func (r *Repo) toDbStatusHistory(statHistory *internal_models.StatusHistory) dbStatusHistory {
	return dbStatusHistory{
		ID:         statHistory.ID,
		SupportID:  statHistory.Support.ID,
		StatusID:   statHistory.Support.Status.ID,
		SelectTime: statHistory.SelectTime,
		ShiftID:    statHistory.Shift.ID,
		Duration:   statHistory.Duration,
	}
}

func toModelStatusHistory(stat *dbStatusHistory) internal_models.StatusHistory {
	return internal_models.StatusHistory{
		ID: stat.ID,
		Support: &internal_models.Support{
			ID: stat.SupportID,
			Status: &internal_models.Status{
				ID: stat.StatusID,
			},
		},
		SelectTime: stat.SelectTime,
		Duration:   stat.Duration,
		Shift: &internal_models.Shift{
			ID: stat.ShiftID,
		},
	}
}

//toDbSupportCard...
func (r *Repo) toDbSupportCard(card *internal_models.Card) dbCard {
	var outCard dbCard = dbCard{
		ID:             card.ID,
		InternalNumber: card.InternalNumber,
		MobileNumber:   card.MobileNumber,
		BirthDate:      card.BirthDate,
		IsSenior:       card.IsSenior,
		Wager:          card.Wager,
		Comment:        card.Comment,
		Color:          card.Color,
	}
	if card.Support != nil {
		outCard.SupportID = card.Support.ID
	}
	if card.Senior != nil {
		outCard.SeniorID = sql.NullInt64{
			Valid: true,
			Int64: int64(card.Senior.ID),
		}
	}
	return outCard
}

func (r *Repo) toModelSupportCard(dbCard *dbCard) *internal_models.Card {
	mCard := internal_models.Card{
		ID:             dbCard.ID,
		Support:        &internal_models.Support{ID: dbCard.ID},
		InternalNumber: dbCard.InternalNumber,
		MobileNumber:   dbCard.MobileNumber,
		BirthDate:      dbCard.BirthDate,
		IsSenior:       dbCard.IsSenior,
		Wager:          dbCard.Wager,
		Comment:        dbCard.Comment,
		Color:          dbCard.Color,
	}
	if dbCard.SeniorID.Valid {
		mCard.Senior = &internal_models.Support{ID: uint64(dbCard.SeniorID.Int64)}
	}
	return &mCard
}