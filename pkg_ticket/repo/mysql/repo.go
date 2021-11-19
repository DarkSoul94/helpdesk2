package mysql

import (
	"database/sql"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewTicketRepo(db *sql.DB) *TicketRepo {
	return &TicketRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *TicketRepo) GetTicketStatuses() ([]*internal_models.TicketStatus, error) {
	var (
		dbList []dbTicketStatus
		mList  []*internal_models.TicketStatus
		query  string
		err    error
	)

	query = `SELECT * FROM ticket_status`

	err = r.db.Select(&dbList, query)
	if err != nil {
		logger.LogError(
			"Failed read ticket status list",
			"pkg_ticket/repo/mysql",
			"",
			err,
		)
		return nil, err
	}

	for _, stat := range dbList {
		mList = append(mList, r.toModelTicketStatus(stat))
	}

	return mList, nil
}

func (r *TicketRepo) Close() error {
	r.db.Close()
	return nil
}
