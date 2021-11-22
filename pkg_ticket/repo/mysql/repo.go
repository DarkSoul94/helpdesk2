package mysql

import (
	"database/sql"
	"fmt"

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

func (r *TicketRepo) CreateTicket(ticket *internal_models.Ticket) (uint64, error) {
	var (
		res   sql.Result
		query string
		err   error
	)

	query = `INSERT INTO tickets SET
	ticket_date = :ticket_date,
	section_id = :section_id,
	ticket_text = :ticket_text,
	ticket_status_id = :ticket_status_id,
	ticket_author_id = :ticket_author_id,
	filial = :filial,
	ip = :ip`

	res, err = r.db.NamedExec(query, r.toDbTicket(ticket))
	if err != nil {
		logger.LogError(
			"Failed create ticket",
			"pkg_ticket/repo/mysql",
			fmt.Sprintf("section id: %d; text: %s; status id: %d; author id: %d; filial: %s; ip: %s", ticket.CatSect.ID, ticket.Text, ticket.Status.ID, ticket.Author.ID, ticket.Filial, ticket.Ip),
			err,
		)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return uint64(id), nil
}

func (r *TicketRepo) Close() error {
	r.db.Close()
	return nil
}
