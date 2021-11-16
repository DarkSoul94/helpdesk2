package mysql

import (
	"database/sql"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

func (r *Repo) toDbSupport(support *internal_models.Support) dbSupport {
	return dbSupport{
		SupportID: support.SupportID,
		StatusID:  support.Status.ID,
		Priority:  support.Priority,
	}
}

func (r *Repo) toModelSupport(dbSupp *dbSupport) *internal_models.Support {
	return &internal_models.Support{
		SupportID: dbSupp.SupportID,
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
		Support:        &models.User{ID: dbCard.SupportID},
		InternalNumber: dbCard.InternalNumber,
		MobileNumber:   dbCard.MobileNumber,
		BirthDate:      dbCard.BirthDate,
		IsSenior:       dbCard.IsSenior,
		Wager:          dbCard.Wager,
		Comment:        dbCard.Comment,
		Color:          dbCard.Color,
	}
	if dbCard.SeniorID.Valid {
		mCard.Senior = &models.User{ID: uint64(dbCard.SeniorID.Int64)}
	}
	return &mCard
}
